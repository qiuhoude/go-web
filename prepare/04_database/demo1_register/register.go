package main

import (
	"database/sql"
	"database/sql/driver"
)

//github.com/go-sql-driver/mysql/driver.go 的伪代码
func init() {
	sql.Register("mysql", &MySQLDriver{})
}

type MySQLDriver struct{}

func (d MySQLDriver) Open(dsn string) (driver.Conn, error) {
	//...
	return nil, nil
}

func main() {

}
