package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var provides = make(map[string]Provider)

type (
	Provider interface {
		SessionInit(sid string) (Session, error)
		SessionRead(sid string) (Session, error)
		SessionDestroy(sid string) error
		SessionGC(maxLifeTime int64)
	}

	Session interface {
		Set(key, value interface{}) error
		Get(key interface{}) interface{}
		Delete(key interface{}) error
		SessionID() string
	}
)

type (
	Manager struct {
		cookieName  string
		lock        sync.Mutex
		maxLifeTime int64 // 最大过期时间
		provider    Provider
	}
)

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		provider:    provider,
	}, nil
}

// 生成sessionId
func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 创建session
func (m *Manager) SessionStart(rw http.ResponseWriter, req *http.Request) Session {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := req.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" { // 获取出错,或者值为nil就重新设置cookie
		sid := m.sessionId()
		cookie := http.Cookie{
			Name:  m.cookieName,
			Value: url.QueryEscape(sid),
			Path:  "/",
			//MaxAge:   int(m.maxLifeTime),
			HttpOnly: true, //这个属性是设置是否可通过客户端脚本访问这个设置的cookie
		}
		http.SetCookie(rw, &cookie)
		session, _ := m.provider.SessionInit(sid)
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ := m.provider.SessionRead(sid)
		return session
	}
}

//重置session
func (m *Manager) SessionDestroy(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()
		m.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{
			Name:     m.cookieName,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
		}
		http.SetCookie(rw, &cookie)
	}
}

func (m *Manager) CookieName() string {
	return m.cookieName
}

// 过期回收,每间隔 maxLifeTime 时间就来清理一次session
func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime)*time.Second, func() {
		m.GC()
	})
}

//注册provide, 不能重复注册,在调用的地方 init中执行
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provide
}
