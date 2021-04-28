package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
	_    [0]byte
}

// 禁止 Quack 方法的内联编译
//go:noinline
func (c *Cat) Quack() {
	println(c.Name + " meow")
}

// 得到汇编代码的方式
// go tool compile -S -l -N main.go > s.txt;
// go tool compile -N -l main.go;
// go build -gcflags -S main.go
func main() {
	var c Duck = &Cat{Name: "draven"}
	switch c.(type) {
	case *Cat:
		cat := c.(*Cat)
		cat.Quack()
	}
}
