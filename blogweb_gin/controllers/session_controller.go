// 用于获取session，查看用户是否登录
package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"net/http"
)

const (
	SessionKey      = "loginSession" // 浏览器 cookie 的名字
	SessionStoreKey = "loginuser"
)

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get(SessionStoreKey)
	logs.Info.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

//退出
func ExitGet(c *gin.Context) {
	session := sessions.Default(c)
	//清除该用户登录状态的数据
	session.Delete(SessionStoreKey)
	session.Save()
	session.Clear()

	logs.Info.Println("delete session...", session.Get(SessionStoreKey))
	c.Redirect(http.StatusMovedPermanently, "/")
}

func IsLoginMiddle(c *gin.Context) {
	if c.Request.Method == http.MethodGet && GetSession(c) {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	c.Next()
}

// 需要登陆才能操作的中间件
func NeedLoginMiddle(c *gin.Context) {
	if !GetSession(c) {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.Next()
}
