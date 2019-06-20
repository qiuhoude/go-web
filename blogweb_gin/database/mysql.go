package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	username = "root"
	pwd      = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbname   = "go_test"
)

var DB *gorm.DB

func InitMsql() {
	// 修改默认前缀
	/*
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "s_"+defaultTableName
		}
	*/
	if DB == nil {
		dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, pwd, ip, port, dbname)
		var err error
		DB, err = gorm.Open("mysql", dbUrl)
		if err != nil {
			panic(err)
		}
		DB.LogMode(true)

		//如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
		//db.SingularTable(true)
		//DB.Set("gorm:table_options", "ENGINE=InnoDB")
		// migrate 仅支持创建表，没有的字段和没有索引,它并不支持改变已有的字段类型或删除未被使用的字段
		//DB.AutoMigrate(&models.User{})

		fmt.Println("InitMysql...")
	}
}
