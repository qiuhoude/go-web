package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"log"
	"net/http"
	"strconv"
)

func DeleteArticleGet(c *gin.Context) {
	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	logs.Info.Println("删除 id:", id)

	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}
	//c.JSON(http.StatusOK, gin.H{"IsLogin": islogin})
	c.Redirect(http.StatusMovedPermanently, "/")
}
