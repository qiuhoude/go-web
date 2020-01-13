package _2_viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

func init() {
	viper.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

func main() {
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			//viper配置发生变化了 执行响应的操作
			fmt.Println("Config file changed:", e.Name, e.Op)
			fmt.Println("update", viper.GetString(`app.name`))
		})
	}()

	fmt.Println("获取配置文件的string", viper.GetString(`app.name`))
	fmt.Println("获取配置文件的string", viper.GetInt(`app.foo`))
	fmt.Println("获取配置文件的string", viper.GetBool(`app.bar`))
	fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit // 没有收到中断信号前,会阻塞在此处
}
