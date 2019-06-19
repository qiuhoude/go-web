package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
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

	router.Run(":8000")
}
