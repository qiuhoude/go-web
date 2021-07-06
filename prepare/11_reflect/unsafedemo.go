package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/11_reflect/rtype"
	"reflect"
	"unsafe"
)

func main() {
	//sizeof()
	//alignof()
	//offsetof()
	//offsetof2()
	//unsafePoint()
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(int(0))))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(int(0))))
	fmt.Println(Cap, cap(s)) // 20 20
}

func unsafePoint() {
	i := 10
	ip := &i
	// 指针转换之间强转不能进行强转,只能通过 unsafe.Pointer
	/*
		1. 任何指针都可以转换为unsafe.Pointer
		2. unsafe.Pointer可以转换为任何指针
		3. uintptr可以转换为unsafe.Pointer
		4. unsafe.Pointer可以转换为uintptr
	*/
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)

	u := new(rtype.User)
	//u :=&rtype.User{Age:ip}
	fmt.Println(*u)
	u.Age = ip

	pName := (*string)(unsafe.Pointer(u)) // 将结构体指针 强转成string指针, 可以强转的原因是, user首字段是string类型
	*pName = "qiu"

	pLove := (*interface{})(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Love)))
	*pLove = "name"

	// 操作lv私有字段
	plv := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Love) + uintptr(unsafe.Sizeof(u.Love))))
	*plv = 14

	fmt.Println(u)
	(*u).Print("pName--")

}

// Sizeof 类型大小
func sizeof() {
	fmt.Println("---------------------Sizeof---------------------")
	fmt.Println("bool ", unsafe.Sizeof(true))
	fmt.Println("int8 ", unsafe.Sizeof(int8(0)))
	fmt.Println("int16 ", unsafe.Sizeof(int16(10)))
	fmt.Println("int32 ", unsafe.Sizeof(int32(10000000)))
	fmt.Println("int64 ", unsafe.Sizeof(int64(10000000000000)))
	fmt.Println("int ", unsafe.Sizeof(int(10000000000000000)))
}

//Alignof 对齐值, 对齐值一般是2^n,最大不会超过8
func alignof() {
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32

	var s string

	var m map[string]string

	var p *int32
	fmt.Println("---------------------Alignof---------------------")
	fmt.Println("bool ", unsafe.Alignof(b))
	fmt.Println("int8 ", unsafe.Alignof(i8))
	fmt.Println("int16 ", unsafe.Alignof(i16))
	fmt.Println("int64 ", unsafe.Alignof(i64))
	fmt.Println("float32 ", unsafe.Alignof(f32))
	fmt.Println("string ", unsafe.Alignof(s))
	fmt.Println("map[string]string ", unsafe.Alignof(m))
	fmt.Println("*int32 ", unsafe.Alignof(p))
}

//Offsetof  函数只适用于struct结构体中的字段相对于结构体的内存位置偏移量。结构体的第一个字段的偏移量都是0
func offsetof() {
	var u1 rtype.User1
	fmt.Println("---------------------Offsetof---------------------")
	fmt.Println("byte ", unsafe.Offsetof(u1.B))
	// unsafe.Offsetof(u1.B) 与 reflect.TypeOf(u1).FieldByName("B").Offset等价
	if field, ok := reflect.TypeOf(u1).FieldByName("B"); ok {
		fmt.Println("reflect byte ", field.Offset)
	}
	fmt.Println("int32 ", unsafe.Offsetof(u1.I32))
	fmt.Println("int64 ", unsafe.Offsetof(u1.I64))
	fmt.Println("u1 sizeof ", unsafe.Sizeof(u1)) //排方式不同对齐方式也不同

}

func offsetof2() {
	var u1 rtype.User
	u1.Love = true
	u1.Print("prifix")
	fmt.Println("---------------------Offsetof2---------------------")
	fmt.Println("Name ", unsafe.Offsetof(u1.Name))
	fmt.Println("Age ", unsafe.Offsetof(u1.Age))
	fmt.Println("Love ", unsafe.Offsetof(u1.Love))
	fmt.Println("size Love ", unsafe.Sizeof(u1.Love))
	fmt.Println("u1 sizeof ", unsafe.Sizeof(u1))

}
