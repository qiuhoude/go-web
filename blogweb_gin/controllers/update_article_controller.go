package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"net/http"
	"strconv"
)

func UpdateArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	c.HTML(http.StatusOK, "write_article.html",
		gin.H{"IsLogin": islogin, "Title": art.Title, "Tags": art.Tags, "Short": art.Short, "Content": art.Content, "Id": art.Id})
}

func UpdateArticlePost(c *gin.Context) {

	//获取表单信息
	idstr := c.PostForm("id")
	id, _ := strconv.Atoi(idstr)
	logs.Info.Println("postid:", id)

	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	logs.Info.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中

	art := models.Article{
		Id:      id,
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
	}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}
	c.JSON(http.StatusOK, response)
}
