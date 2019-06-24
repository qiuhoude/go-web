package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// 静态文件
	//router.StaticFS("/static",http.Dir("static"))
	router.Static("/static", "./static")

	//设置session 中间件
	store := cookie.NewStore([]byte(controllers.SessionStoreKey))
	router.Use(sessions.Sessions(controllers.SessionKey, store))
	{
		//注册
		router.GET("/register", controllers.IsLoginMiddle, controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		//登录
		router.GET("/login", controllers.IsLoginMiddle, controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		// 首页
		router.GET("/", controllers.HomeGet)
		// 退出
		router.GET("/exit", controllers.ExitGet)

		// 路由组
		v1 := router.Group("/article")
		{
			v1.GET("/add", controllers.AddArticleGet)
			v1.POST("/add", controllers.AddArticlePost)
		}
	}
	return router
}
