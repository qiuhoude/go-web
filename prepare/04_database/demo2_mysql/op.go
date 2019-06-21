package demo2_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func open() {
	/*
		连接数据库:func Open(driverName, dataSourceName string) (*DB, error)
		Open打开一个dirverName指定的数据库，dataSourceName指定数据源，
		一般包至少括数据库文件名和（可能的）连接信息。

		driverName: 使用的驱动名. 这个名字其实就是数据库驱动注册到 database/sql 时所使用的名字.
		dataSourceName: 数据库连接信息，这个连接包含了数据库的用户名, 密码, 数据库主机以及需要连接的数据库名等信息.

		drvierName,"mysql"
		dataSourceName,用户名:密码@协议(地址:端口)/数据库?参数=参数值

		sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验, 它仅仅是初始化一个sql.DB对象.
		当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接

		sql.Open返回的sql.DB对象是协程并发安全的.

		sql.DB的设计就是用来作为长连接使用的。不要频繁Open, Close。比较好的做法是，为每个不同的datastore建一个DB对象，保持这些对象Open。
		如果需要短连接，那么把DB作为参数传入function，而不要在function中Open, Close。
	*/
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	fmt.Println("db_err : ", err)
	// 立即验证连接，需要用Ping()方法
	err = db.Ping()
	fmt.Println("db : ", db)
	fmt.Println("ping_err : ", err)
	if err != nil {
		fmt.Println("连接有误。。", err)
		return
	}
	fmt.Println("连接成功。。")
	db.Close()
}

func insert() {
	// 打开数据库，相当于和数据库建立连接：db对象
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 插入语句
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)")
	if err != nil {
		fmt.Println("操作失败。。")
	}
	//补充完整sql语句，并执行
	result, err := stmt.Exec("荣耀", "服务端", "2019-06-04")
	if err != nil {
		fmt.Println("插入数据失败。。", err)
	}
	lastInsertId, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)

	//再次插入数据：
	result, _ = stmt.Exec("ruby", "人事部", "2019-06-04")
	count, _ := result.RowsAffected()
	fmt.Println("影响的行数：", count)

	// 关闭资源
	stmt.Close()
	db.Close()
}

func update() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := db.Exec("update userinfo SET username = ?, departname = ? WHERE uid = ?", "李维民", "公安局", 2)
	if err != nil {
		fmt.Println("更新数据失败。。", err)
	}
	lastInsertId, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)

	db.Close()
}

/**
查询一条
*/
func queryOne() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	row := db.QueryRow("select uid,username,departname,created from userinfo where uid=?", 1)
	var (
		uid                           int
		username, departname, created string
	)
	/*
		   	row：Scan()-->将查询的结果从row取出
			   err对象
			   判断err是否为空，
				   为空，查询有结果，数据可以成功取出,如果是多条数据就是取第1条
				   不为空，没有数据，sql: no rows in result set
	*/
	err = row.Scan(&uid, &username, &departname, &created)
	if err != nil {
		fmt.Println("查询出错 -> ", err)
		return
	}
	fmt.Println(uid, username, departname, created)
}

type User struct {
	uid        int
	username   string
	departname string
	created    string
}

/*
查询多条
*/
func queryMulti() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select uid,username,departname,created from userinfo where uid>?", 0)
	if err != nil {
		fmt.Println("查询出错 -> ", err)
		return
	}
	defer rows.Close()

	fmt.Println(rows.Columns()) //[uid username departname created]

	//创建slice，存入struct，
	var datas []User
	// 操作结果集获取数据
	for rows.Next() {
		var (
			uid                           int
			username, departname, created string
		)
		if err := rows.Scan(&uid, &username, &departname, &created); err != nil {
			fmt.Println("获取失败。。")
		}
		//每读取一行，创建一个user对象，存入datas2中
		user := User{uid, username, departname, created}
		datas = append(datas, user)
	}

	// 打印
	for _, v := range datas {
		fmt.Println(v)
	}
}

func transaction() {
	/*
		事务：
			   4大特性：ACID
			   原子性：
			   一致性：
			   隔离性：
			   永久性：
	*/
	//ruby-->王二狗,2000元

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//开启事务
	tx, _ := db.Begin()
	//提供一组sql操作
	var aff1, aff2 int64 = 0, 0
	result1, _ := tx.Exec("UPDATE account SET money=money-1 WHERE id=?", 1)
	result2, _ := tx.Exec("UPDATE account SET money=money+1 WHERE id=?", 2)
	//fmt.Println(result2)
	if result1 != nil {
		aff1, _ = result1.RowsAffected()
		fmt.Println("aff1:", aff1)
	}
	if result2 != nil {
		aff2, _ = result2.RowsAffected()
		fmt.Println("aff2:", aff2)
	}

	/*
		一个Tx会在整个生命周期中保存一个连接，然后在调用commit()或Rollback()的时候释放掉。
		在调用这几个函数的时候必须十分小心，否则连接会一直被占用直到被垃圾回收
	*/
	if aff1 == 1 && aff2 == 1 {
		//提交事务
		tx.Commit()
		fmt.Println("操作成功。。")
	} else {
		//回滚
		tx.Rollback()
		fmt.Println("操作失败。。。回滚。。")
	}
}

func queryAllDb() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/honor_ini?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	// select table_name from information_schema.tables;
	// show tables;
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println("查询出错 -> ", err)
		return
	}
	//读出查询出的列字段名
	var tableNames []string
	for rows.Next() {
		var tname string
		rows.Scan(&tname)
		tableNames = append(tableNames, tname)
	}
	rows.Close()
	for _, tname := range tableNames {
		queryDataByTable(tname, db)
	}
}

type TableData struct {
	columns *[]string // 表头
	data    *[][]byte // 数据
}

func queryDataByTable(tname string, db *sql.DB) {
	rows, err := db.Query("select * from " + tname)
	if err != nil {
		fmt.Println(tname, "表查询出错 -> ", err.Error())
		return
	}
	defer rows.Close()

	//读出查询出的列字段名
	cols, _ := rows.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	//最后得到的map
	results := make(map[int]map[string]string)
	i := 0
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}

	//查询出来的数组
	for k, v := range results {
		fmt.Println(k, v)
	}
	//fmt.Println(tname + " 表 查询完成")
}

//func formatToLua(data *TableData) string {
//	var b strings.Builder
//	b.WriteString("return {\n")
//	// 拼接头部
//
//
//
//
//	b.WriteString("}")
//}
