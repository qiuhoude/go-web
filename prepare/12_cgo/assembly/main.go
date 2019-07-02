// go 调用汇编
package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/12_cgo/assembly/add"
)

func main() {
	fmt.Println(add.Add(2, 5))
}
