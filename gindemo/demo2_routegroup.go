package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8000")
}

func loginEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func submitEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func readEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}
