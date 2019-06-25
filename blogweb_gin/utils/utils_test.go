package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bmizerany/assert"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"regexp"
	"testing"
	"time"
)

func TestSwitchTimeStampToData(t *testing.T) {
	now := time.Now()
	t.Log(SwitchTimeStampToData(now.Unix()))
	tim, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:30:00", time.Local)

	t.Log(tim.Round(1 * time.Hour)) //接近,四舍五入的方式
	t.Log(tim.Truncate(1 * time.Hour))
	t.Log(int(now.Weekday()))
	t.Log(int(now.Month()))
	t.Log(now.Day()) // dayofmouth
	t.Log(now.YearDay())
	t.Log(now.ISOWeek())
}

func TestMdToHtml(t *testing.T) {
	mdStr := `
	## 一、russross/blackfriday包`
	markdown := blackfriday.MarkdownCommon([]byte(mdStr))
	doc, _ := goquery.NewDocumentFromReader(bufio.NewReader(bytes.NewReader(markdown)))
	htmlString, _ := doc.Html()
	t.Log(htmlString)
}

func TestGoQuery(t *testing.T) {

	doc, err := goquery.NewDocument("http://studygolang.com/topics") // goquery 用于解析html
	if err != nil {
		t.Fatal(err)
	}

	doc.Find(".topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		//Find(".title a")与Find(".title").Find("a")一样
		fmt.Println("第", i+1, "个帖子的标题：", title)
		//ret,_ := contentSelection.Html()
		//fmt.Printf("\n\n\n%v", ret)
		//fmt.Println(contentSelection.Text())
	})
	//最终输出为 html 文档：
	//new, err := doc.Html()
}

// 语法高亮
func TestSyntaxhighlight(t *testing.T) {
	src := []byte(`
/* hello, world! */
var a = 3;

// b is a cool function
function b() {
  return 7;
}`)

	highlighted, err := syntaxhighlight.AsHTML(src)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(highlighted))
}

func TestRe(t *testing.T) {
	re := regexp.MustCompile(`\?`)
	s := re.ReplaceAllString("inset into album(filepath,filename,status,createtime)values(?,?,?,?)", "%v")
	assert.Equal(t, `inset into album(filepath,filename,status,createtime)values(%v,%v,%v,%v)`, s)
}
