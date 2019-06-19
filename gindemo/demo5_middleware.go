package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 有点类似于netty中的 pipelie, HandlerFunc就类似于ChannelHandler
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("before middleware")
		//设置request变量到Context的Key中,通过Get等函数可以取得
		c.Set("request", "client_request")
		//Request
		c.Next()
		//response
		// 这个c.Write是ResponseWriter,我们可以获得状态等信息
		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t) // t2 - t1
		fmt.Println("time", t2)
	}
}

// 模拟私有数据
var secrets = gin.H{
	"houdeqiu": gin.H{"email": "hanru@163.com", "phone": "123433"},
	"enqiqiu":  gin.H{"email": "wangergou@example.com", "phone": "666"},
}

func main() {
	r := gin.Default()

	// 使用中间件
	r.Use(MiddleWare())
	r.GET("/middleware", func(c *gin.Context) {
		//获取gin上下文中的变量
		request := c.MustGet("request").(string)
		req, _ := c.Get("request")
		fmt.Println("request:", request)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
			"request":         req,
		})
		fmt.Println("访问 /middleware")
	})

	// 单个路由上加上中间件
	r.GET("/before", MiddleWare(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	authorize := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"houdeqiu": "123",
		"enqiqiu":  "456",
	}))

	// 定义路由
	authorize.GET("/secrets", func(c *gin.Context) {
		// 获取提交的用户名（AuthUserKey）
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.Run(":8000")
}
