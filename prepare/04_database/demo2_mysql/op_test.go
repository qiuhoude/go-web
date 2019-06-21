package demo2_mysql

import (
	"errors"
	"os"
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
	Title []string
	Data  interface{}
}

func TestTemplate(t *testing.T) {
	const mm = `
return {
	title={ {{- range $i,$d:= .Title -}} {{- if (eq $i 0) -}},"{{- $d -}}" {{- else -}}"{{- $d -}}" {{ end -}} {{- end -}} },
		records={
				
		}
}
`
	data := lua{
		Title: []string{"id", "activityId", "time"},
	}

	var fmap = template.FuncMap{
		"sequence": sequenceFunc,
		"cycle":    cycleFunc,
	}

	tp := template.New("lua").Funcs(fmap)
	//tp.Delims("{{", "}}")
	tp = template.Must(tp.Parse(mm))

	tp.Execute(os.Stdout, data)
}

type generator struct {
	ss []string
	i  int
	f  func(s []string, i int) string
}

func sequenceGen(ss []string, i int) string {
	if i >= len(ss) {
		return ss[len(ss)-1]
	}
	return ss[i]
}

func cycleGen(ss []string, i int) string {
	return ss[i%len(ss)]
}

func sequenceFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("sequence must have at least one element")
	}
	return &generator{ss, 0, sequenceGen}, nil
}

func cycleFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("cycle must have at least one element")
	}
	return &generator{ss, 0, cycleGen}, nil
}
