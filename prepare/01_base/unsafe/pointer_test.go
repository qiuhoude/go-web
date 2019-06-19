package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func covPointer() {
	// Go语言是不允许两个指针类型进行转换的。

	i := 9
	ip := &i

	//var fp *float64 =(*float64)(ip) // 不让转换

	var fp *float64 = (*float64)(unsafe.Pointer(ip))

	fmt.Printf("%T %v\n", fp, fp)
	*fp = *fp * 0.5
	fmt.Printf("%T %g\n", *fp, *fp)
	fmt.Println(i)
}

func TestCovPointer(t *testing.T) {
	covPointer()
}

type user struct {
	name string
	age  int
}

func TestUintptr(t *testing.T) {
	u := new(user)
	fmt.Println(*u)

	// 因为name 是第一个字段,位置是从0开始,所以可以这样获取
	pName := (*string)(unsafe.Pointer(u))
	*pName = "Alex"

	// 要找到age 必须找到,age的offset
	fmt.Println("u.age offset ", unsafe.Offsetof(u.age))

	//up := uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)
	//fmt.Printf("up : %T %v \n ", up, up)
	//pAge := (*int)(unsafe.Pointer(up))

	// 表达式非常长，但是也千万不要把他们分段, 这里会牵涉到GC，如果我们的这些临时变量被GC,就不知道是哪块内存了，会引起莫名其妙的问题
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))

	*pAge = 32
	fmt.Println(*u)
}
