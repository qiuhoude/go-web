package unsafe

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestSizeOf(t *testing.T) {
	//Sizeof 函数可以返回一个类型所占用的内存大小，这个大小只有类型有关，.
	// 和类型对应的变量存储的内容大小无关，比如bool型占用一个字节、int8也占用一个字节
	t.Log(unsafe.Sizeof(true))
	t.Log(unsafe.Sizeof(int8(0)))
	t.Log(unsafe.Sizeof(int16(10)))
	t.Log(unsafe.Sizeof(int32(10000000)))
	t.Log(unsafe.Sizeof(int64(10000000000000)))
	t.Log(unsafe.Sizeof(int(10000000000000000)))
}

func TestAlignof(t *testing.T) {
	//var b bool
	//var i8 int8
	//var i16 int16
	//var i32 int32
	//var i64 int64
	//
	//var f32 float32
	//
	//var s string
	//
	//var m map[string]string
	//
	//var p *int32
	////Alignof 返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数。
	//// 对齐值是一个和内存对齐有关的值，合理的内存对齐可以提高内存读写的性能
	//// unsafe.Alignof(x) 等价于 reflect.TypeOf(x).Align()
	//fmt.Printf("%T %v\n", b, unsafe.Alignof(b))
	//fmt.Printf("%T %v\n", i8, unsafe.Alignof(i8))
	//fmt.Printf("%T %v\n", i16, unsafe.Alignof(i16))
	//fmt.Printf("%T %v\n", i32, unsafe.Alignof(i32))
	//fmt.Printf("%T %v\n", i64, unsafe.Alignof(i64))
	//fmt.Printf("%T %v\n", f32, unsafe.Alignof(f32))
	//fmt.Printf("%T %v\n", s, unsafe.Alignof(s))
	//fmt.Printf("%T %v\n", m, unsafe.Alignof(m))
	//fmt.Printf("%T %v\n", p, unsafe.Alignof(p))
	//fmt.Printf("%T %v\n", p, reflect.TypeOf(p).Align())
	//fmt.Printf("%T %v\n", p, unsafe.Alignof(p))

	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	// 整点（向下取整）
	fmt.Println(ti.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(ti.Round(1 * time.Hour))

	// 整分（向下取整）
	fmt.Println(ti.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(ti.Round(1 * time.Minute))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", ti.Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t2)
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	ticker := time.NewTicker(5 * time.Second)
	<-ticker.C

}

type user1 struct {
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}

func TestOffsetof(t *testing.T) {
	// Offsetof函数只适用于struct结构体中的字段相对于结构体的内存位置偏移量。结构体的第一个字段的偏移量都是0.
	// unsafe.Offsetof(u1.i)等价于reflect.TypeOf(u1).Field(i).Offset
	var u1 user2

	fmt.Println(unsafe.Offsetof(u1.b))
	fmt.Println(unsafe.Offsetof(u1.i))
	fmt.Println(unsafe.Offsetof(u1.j))
}

func TestSzieofStruct(t *testing.T) {
	// 因为有对其方式,每个结构体的大小就不同
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6

	fmt.Println("u1 size is ", unsafe.Sizeof(u1))
	fmt.Println("u2 size is ", unsafe.Sizeof(u2))
	fmt.Println("u3 size is ", unsafe.Sizeof(u3))
	fmt.Println("u4 size is ", unsafe.Sizeof(u4))
	fmt.Println("u5 size is ", unsafe.Sizeof(u5))
	fmt.Println("u6 size is ", unsafe.Sizeof(u6))

}
