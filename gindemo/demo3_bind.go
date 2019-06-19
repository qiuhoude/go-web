package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	// json绑定
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
		//c.BindJSON(json) 这种方式 gin框架帮忙处理错误 c.AbortWithError (400) .SetType (ErrorTypeBind)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "hanru" || json.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		//curl -v -X POST http://127.0.0.1:8000/loginJSON -H 'content-type:application/json' -d '{"user":"hanru","password":"hanru123"}'

	})

	// form绑定
	router.POST("/loginFrom", func(c *gin.Context) {
		var form Login
		//方法一：对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
		// 根据请求头中 content-type 自动推断.

		//方法二: 使用BindWith函数,如果你明确知道数据的类型,可以显式声明来绑定多媒体表单： c.BindWith(&form, binding.Form)或者使用自动推断:
		//err := c.BindWith(&form, binding.Form) err := c.BindWith(&form, binding.JSON)

		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "hanru" || form.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// url绑定
	router.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.BindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"username": login.User, "password": login.Password})
	})

	router.Run(":8000")
}
