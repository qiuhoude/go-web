package main

import (
	"context"
	"github.com/qiuhoude/go-web/blogweb_gin/database"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	_ "github.com/qiuhoude/go-web/blogweb_gin/logs"
	"github.com/qiuhoude/go-web/blogweb_gin/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	database.InitMsql()
	router := routers.InitRouter()

	//router.Run(":8000")

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		// 开启监听服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	// 提示中断
	signal.Notify(quit, os.Interrupt)
	<-quit // 没有收到中断信号前,会阻塞在此处
	logs.Error.Println("Shutdown Server ...")

	// 超时的ctx ,超时会自己执行cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logs.Error.Println("Server exiting")
}
