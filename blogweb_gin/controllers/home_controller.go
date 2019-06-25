package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"net/http"
	"strconv"
)

//主页
//可以通过翻页来获取该网页，也可以通过tag标签获取
//传page参数代表翻页，传tag参数代表标签
//首先判断page有值那么就是翻页，否则判断tag有值就是标签，否则就是默认的第一页
func HomeGet(c *gin.Context) {
	//获取session，判断用户是否登录
	islogin := GetSession(c)

	tag := c.Query("tag")
	logs.Info.Println("tag:", tag)
	page, _ := strconv.Atoi(c.Query("page"))

	var hasFooter bool
	var artList []models.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		hasFooter = false
	} else {
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticleWithPage(page)
		hasFooter = true
	}

	homeFooterPageCode := models.ConfigHomeFooterPageCode(page)
	html := models.MakeHomeBlocks(artList, islogin)

	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin, "Content": html, "HasFooter": hasFooter, "PageCode": homeFooterPageCode})
}
