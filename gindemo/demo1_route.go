package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Ketty")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"name": name})
		c.String(http.StatusOK, name)
	})
	// url参数
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") // 可设置默认值
		//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
		c.String(http.StatusOK, name)
	})

	// * 匹配后面的所有
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	//  表单参数
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		// c.PostForm解析的是x-www-form-urlencoded或from-data的参数。
		username := c.PostForm("username")
		password := c.PostForm("password")

		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s,hobby is %v", type1, username, password, hobbys))
	})

	router.GET("/form", func(c *gin.Context) {
		t := template.Must(template.ParseFiles("static/login.html"))
		t.Execute(c.Writer, nil)
	})

	router.GET("/upload", func(c *gin.Context) {
		t := template.Must(template.ParseFiles("static/file.html"))
		t.Execute(c.Writer, nil)
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single fileHeader
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusOK, "名称错误")
			return
		}
		log.Println(fileHeader.Filename)

		// Upload the fileHeader to specific dst.
		c.SaveUploadedFile(fileHeader, fileHeader.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileHeader.Filename))
		//file, _ := os.Open(fileHeader.Filename)
		//io.Copy(c.Writer, file)
	})

	// 自定义服务器
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
