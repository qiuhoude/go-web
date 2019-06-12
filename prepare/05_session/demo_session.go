package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/05_session/session"
	_ "github.com/qiuhoude/go-web/prepare/05_session/session/memory"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}

// session管理
var globalSessions *session.Manager

const (
	keyCount      = "countnum"
	keyToken      = "token"
	keyCreateTime = "sessionKeyCreateTime"
)

func init() {
	globalSessions, _ = session.NewManager("memory", "goSessionId", 3600)
	fmt.Println("globalSessions,", globalSessions)
	go globalSessions.GC()
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(globalSessions.CookieName())
	if cookie != nil {
		fmt.Fprintf(w, "已经登录 cookie is %v", cookie)
	} else {
		io.WriteString(w, "hi index")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login....")
	sess := globalSessions.SessionStart(w, r) // 获取session
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", http.StatusFound) // 重定向 到 /
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	fmt.Println("count.......")
	sess := globalSessions.SessionStart(w, r) // 获取session
	fmt.Println("sessID:", sess.SessionID())

	// session劫持防范的方式
	// 方案1 cookie的httponly为true; 在每个请求里面加上隐藏的token
	//h := md5.New()
	//salt := "iamsalt"
	//io.WriteString(h, salt+time.Now().String())
	//token := fmt.Sprintf("%x", h.Sum(nil))
	//if r.Form["token"][0] != token {
	//	//提示登录
	//}
	//sess.Set(sessionKeyToken, token)

	// 方案2
	// 给session额外设置一个创建时间的值，一旦过了一定的时间，我们销毁这个sessionID，重新生成新的session
	createTime := sess.Get(keyCreateTime)
	if createTime == nil {
		sess.Set(keyCreateTime, time.Now().Unix())
	} else if createTime.(int64)+60 < time.Now().Unix() { //过期销毁,并进行重新创建sessionId
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}

	ct := sess.Get(keyCount)
	fmt.Println("ct:", ct)
	if ct == nil {
		sess.Set(keyCount, 1)
	} else {
		sess.Set(keyCount, ct.(int)+1)
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get(keyCount))
}
