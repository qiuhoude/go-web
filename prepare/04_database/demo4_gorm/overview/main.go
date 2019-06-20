package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	Uid        uint   `gorm:"primary_key;column:uid"`
	Username   string `gorm:"column:username"`
	Departname string `gorm:"column:departname"`
	Created    string `gorm:"column:created"`
}

func (UserInfo) TableName() string {
	return "userinfo"
}

type User struct {
	gorm.Model
	Name     string
	Age      int `gorm:"default:18"`
	Birthday time.Time
}

//func (user *User) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("ID", 1)
//	return nil
//}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// 开启 Logger, 以展示详细的日志
	db.LogMode(true)
	//db.SetLogger(log.New(os.Stdout, "【】", 0))

	//fmt.Println(db.HasTable(&UserInfo{}))
	// 自动迁移模式 ,给给
	//db.AutoMigrate(&UserInfo{})

	//createTable(db)
	//create(db)
	query(db)
}

func query(db *gorm.DB) {
	//var listU []UserInfo
	//db.Find(&listU)
	//fmt.Println(listU)

	u := UserInfo{}
	db.Where("username= ?", "李维民").Find(&u)
	fmt.Println(u)
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	//db.Table("user").CreateTable(&User{})
}

func create(db *gorm.DB) {
	user := User{Name: "Jinzhu", Age: 20, Birthday: time.Now()}
	record1 := db.NewRecord(user) // => 返回 `true` ，因为主键为空
	fmt.Println("record1=", record1, user.ID)
	db.Create(&user)
	record2 := db.NewRecord(user) // => 在 `user` 之后创建返回 `false`
	fmt.Println("record2=", record2, user.ID)
}
