package models

import (
	"fmt"
	"github.com/qiuhoude/go-web/blogweb_gin/config"
	"github.com/qiuhoude/go-web/blogweb_gin/database"
	"github.com/qiuhoude/go-web/blogweb_gin/logs"
	"log"
	"strconv"
	"sync"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}

//---------添加文章-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	logs.Info.Println("---------->page", page)
	//从配置文件中获取每页的文章数量
	return QueryArticleWithPage(page, config.NUM)
}

/*
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := database.QueryDb(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

//------翻页------

//存储表的行数， 当文章新增或者删除时需要更新这个值
var artcileRowsNum int = -1
var rwLock sync.RWMutex

//查询文章的总条数
func QueryArticleRowNum() int {
	row := database.QueryRowDb("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if artcileRowsNum == -1 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//设置页数
func SetArticleRowsNum() {
	rwLock.Lock()
	defer rwLock.Unlock()
	artcileRowsNum = QueryArticleRowNum()
}

//----------查询文章-------------

func QueryArticleWithId(id int) Article {
	row := database.QueryRowDb("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

//----------修改数据----------

func UpdateArticle(article Article) (int64, error) {
	//数据库操作
	return database.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//----------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithArtId(artID int) (int64, error) {
	return database.ModifyDB("delete from article where id=?", artID)
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := database.QueryDb(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]Article, error) {

	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	logs.Info.Println(sql)
	return QueryArticlesWithCon(sql)
}
