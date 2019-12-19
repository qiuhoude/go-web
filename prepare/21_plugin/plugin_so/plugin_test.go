package main

import (
	"log"
	"time"
)

/*
编写plugin插件要点

包名称必须是main
没有main函数
必须有可以导出(访问)的变量或者方法
*/
func init() {
	log.Println("plugin init function called")
}

type PluginTest string

func (p PluginTest) GetNowTime() string {
	ret := time.Now().Format(time.RFC822)
	log.Println("call plugin >>> ", ret)
	return ret
}

//go build -buildmode=plugin -o=plugin_test.so plugin_test.go

// 到处可用的
var PluginT = PluginTest(time.Now().Format(time.RFC3339))
