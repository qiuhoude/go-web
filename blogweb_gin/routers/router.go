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
			//添加文章
			v1.GET("/add", controllers.NeedLoginMiddle, controllers.AddArticleGet)
			v1.POST("/add", controllers.NeedLoginMiddle, controllers.AddArticlePost)

			// 显示文章
			v1.GET("/show/:id", controllers.ShowArticleGet)

			//更新文章
			v1.GET("/update", controllers.NeedLoginMiddle, controllers.UpdateArticleGet)
			v1.POST("/update", controllers.NeedLoginMiddle, controllers.UpdateArticlePost)

			//删除文章
			v1.GET("/delete", controllers.NeedLoginMiddle, controllers.DeleteArticleGet)
		}
		router.GET("/tags", controllers.TagsGet)

		router.GET("/album", controllers.AlbumGet)
		// 文件上传
		router.POST("/upload", controllers.UpladPost)
		//关于我
		router.GET("/aboutme", controllers.AboutMeGet)
	}
	return router
}
