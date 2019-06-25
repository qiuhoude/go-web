package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"net/http"
)

func TagsGet(c *gin.Context) {

	islogin := GetSession(c)

	tags := models.QueryArticleWithParam("tags")
	tagmap := models.HandleTagsListData(tags)

	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": tagmap, "IsLogin": islogin})
}
