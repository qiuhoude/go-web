package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//router.StaticFS("/static",http.Dir("static"))
	router.Static("/static", "./static")
	router.GET("/register", controllers.RegisterGet)
	return router
}
