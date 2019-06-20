package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qiuhoude/go-web/blogweb_gin/database"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(64)"`
	Password string `gorm:"column:password;type:varchar(64)"`
}

//设置 `User` 的表名为 `profiles`
func (User) TableName() string {
	return "users"
}

func InsertUser(u *User) {
	database.DB.NewRecord(u)
	database.DB.Create(u)
}
