package demo2_mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User2 struct {
	uid        int
	username   string
	departname string
	created    string
}

// sql的扩展包
func run() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	defer db.Close()
	var Db *sqlx.DB
	Db = db
	var users []User2
	err = Db.Select(&users, "SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("this is Select res:%v\n", users)
	//var user User
	//err1 := Db.Get(&user, "SELECT uid,username,departname,created FROM userinfo where uid = ?", 1)
	//if err1 != nil {
	//	fmt.Println("GET error :", err1)
	//} else {
	//	fmt.Printf("this is GET res:%v", user)
	//}
}
