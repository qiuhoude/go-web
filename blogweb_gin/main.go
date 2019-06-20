package main

import (
	"github.com/qiuhoude/go-web/blogweb_gin/database"
	"github.com/qiuhoude/go-web/blogweb_gin/routers"
)

func main() {
	database.InitMsql()
	router := routers.InitRouter()
	router.Run(":8000")
}
