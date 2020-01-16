package main

import (
	"github.com/sirupsen/logrus"

	"os"
)

var Logger = logrus.New()

func init() {
	// 设置日志格式为json格式
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Logger.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	Logger.SetLevel(logrus.WarnLevel)

}

func main() {
	Logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	Logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	Logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

}
