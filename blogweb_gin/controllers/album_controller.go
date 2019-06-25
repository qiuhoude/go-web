package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/models"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func AlbumGet(c *gin.Context) {
	islogin := GetSession(c)
	albums, _ := models.FindAllAlbums()

	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": islogin, "Album": albums})
}

func UpladPost(c *gin.Context) {
	fileHeader, err := c.FormFile("upload")
	if err != nil {
		responseErr(c, err)
		return
	}
	logs.Info.Println("name:", fileHeader.Filename, fileHeader.Size)

	fileExt := filepath.Ext(fileHeader.Filename)
	logs.Info.Println("ext:", filepath.Ext(fileHeader.Filename)) // 扩展名
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}

	now := time.Now()
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm) //创建对应的文件夹
	if err != nil {
		responseErr(c, err)
		return
	}
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	c.SaveUploadedFile(fileHeader, filePathStr)

	filePathStrUnix := filepath.ToSlash(filePathStr)
	if fileType == "img" {
		// 如果是图片存到数据库
		album := models.Album{0, filePathStrUnix, fileName, 0, timeStamp} //
		models.InsertAlbum(album)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
}
