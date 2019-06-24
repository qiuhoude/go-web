package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
)

const (
	username = "root"
	pwd      = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbname   = "go_test"
)

var db *sql.DB

func InitMsql() {
	// 修改默认前缀
	/*
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "s_"+defaultTableName
		}
	*/
	fmt.Println("InitMysql...")
	if db == nil {
		dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, pwd, ip, port, dbname)
		var err error
		db, err = sql.Open("mysql", dbUrl)
		if err != nil {
			panic(err)
		}
		// 初始化数据库
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	logs.Info.Println(sql)
	result, err := db.Exec(sql, args...)
	if err != nil {
		logs.Error.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		logs.Error.Println(err)
		return 0, err
	}
	return count, err
}

// 查询单条
func QueryRowDb(sql string, args ...interface{}) *sql.Row {
	return db.QueryRow(sql, args...)
}

// 查询多条
func QueryDb(sql string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(sql, args...)
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	ModifyDB(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

// 图片
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
