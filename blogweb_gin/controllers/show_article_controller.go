package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"github.com/qiuhoude/go-web/blogweb_gin/utils"
	"net/http"
	"strconv"
)

// 显示文章
func ShowArticleGet(c *gin.Context) {
	islogin := GetSession(c)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	logs.Info.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	contentHTML := utils.SwitchMarkdownToHtml(art.Content)

	//渲染HTML
	c.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Content": contentHTML})
}
