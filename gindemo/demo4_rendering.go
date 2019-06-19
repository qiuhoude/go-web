// json/xml/yaml渲染
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("someJSON", func(c *gin.Context) {
		// gin.H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{"message": "hi", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "qiu"
		msg.Message = "hi"
		msg.Number = 11
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"user": "qiu", "message": "hi", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"user": "qiu", "message": "hi", "status": http.StatusOK})
	})

	// proto解析
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"

		data := protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, &data)
	})

	// html模板

	// 加载全局模板
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/login.html","templates/demo.html")

	r.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index1.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// 不同文件夹下模板名字可以相同，此时需要 LoadHTMLGlob() 加载两层模板路径。
	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.tmpl", gin.H{
			"title": "Posts",
		})
		c.HTML(http.StatusOK, "index1.tmpl", gin.H{
			"title": "Users",
		})
	})

	//r.GET("/JSONP?callback=x", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"foo": "bar",
	//	}
	//	// callback 是 x
	//	// 将输出：x({\"foo\":\"bar\"})
	//	c.JSONP(http.StatusOK, data)
	//})

	//显示当前文件夹下的所有文件/或者指定文件
	// 定义多文件的路径,使用的是系统的路径(绝对,相对地址都可以)
	r.StaticFS("/fs", http.Dir("."))
	r.StaticFS("/360", http.Dir("/360Downloads"))

	//Static提供给定文件系统根目录中的文件。
	//router.Static("/files", "/bin")
	r.StaticFile("/image", "./assets/aa.jpeg")

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// 异步
	r.GET("/long_async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			// 注意使用只读上下文
			log.Println("Done! in path " + cCp.Request.URL.Path)
			c.String(http.StatusMovedPermanently, "睡醒了")
		}()

	})
	// 同步
	r.GET("/long_sync", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		time.Sleep(3 * time.Second)
		// 注意使用只读上下文
		log.Println("Done! in path " + cCp.Request.URL.Path)
		c.String(http.StatusMovedPermanently, "睡醒了")

	})
	r.Run(":8000")

}
