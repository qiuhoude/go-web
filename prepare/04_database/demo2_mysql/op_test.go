package demo2_mysql

import (
	"bytes"
	"os"
	"strconv"
	"testing"
	"text/template"
)

func TestOpen(t *testing.T) {
	open()
}

func TestInsert(t *testing.T) {
	insert()
}

func TestUpdate(t *testing.T) {
	update()
}

func TestQueryOne(t *testing.T) {
	queryOne()
}
func TestQueryMulti(t *testing.T) {
	queryMulti()
}

func TestTransaction(t *testing.T) {
	transaction()
}
func TestQueryAllDb(t *testing.T) {
	queryAllDb()

}

type lua struct {
	Title []interface{}
	Data  interface{}
}

func TestTemplate(t *testing.T) {
	const mm = `
	return {
		title={ {{strupper .Title}} },
			records={
				
			}
	}
`
	d := []string{"id", "activityId", "time", "xixi"}
	var dataArr []interface{}
	dataArr = append(dataArr, 66)
	for _, v := range d {
		dataArr = append(dataArr, v)
	}
	dataArr = append(dataArr, 66)
	dataArr = append(dataArr, 45)

	data := lua{
		Title: dataArr,
	}

	var fmap = template.FuncMap{
		"strupper": func(ss []interface{}) string {
			var b bytes.Buffer
			l := len(ss)
			for i, s := range ss {
				switch s.(type) {
				case string:
					if i == l-1 {
						b.WriteString("\"")
						b.WriteString(s.(string))
						b.WriteString("\"")
					} else {
						b.WriteString("\"")
						b.WriteString(s.(string))
						b.WriteString("\",")
					}
				case int:
					if i == l-1 {
						b.WriteString(strconv.Itoa(s.(int)))
					} else {
						b.WriteString(strconv.Itoa(s.(int)) + ",")
					}
				}
			}
			return b.String()
		},
	}
	tp := template.New("").Funcs(fmap)
	//tp.Delims("{{", "}}")
	tp = template.Must(tp.Parse(mm))

	tp.Execute(os.Stdout, data)
}
