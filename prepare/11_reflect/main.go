package main

import (
	"errors"
	"fmt"
	"github.com/qiuhoude/go-web/prepare/11_reflect/rtype"
	"reflect"
	"unsafe"
)

func main() {
	age := 10
	pAge := &age
	u := rtype.User{
		Name: "qiu",
		Age:  pAge,
	}
	u.SetLv(100)

	t := reflect.TypeOf(u)
	fmt.Println("t: ", t)

	// value会包含
	v := reflect.ValueOf(u)
	fmt.Println("v:", v)

	vt := v.Type()
	fmt.Println("vt:", vt)

	// 将reflect.Value 转成原始的对象
	u2 := v.Interface().(rtype.User)
	fmt.Println("u2: ", u2)

	//底层类型 比如 interface struct int64 ...
	tkind := t.Kind()
	fmt.Println("tkind: ", tkind)

	// 遍历所有字段, 没有导出的字段也会有
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("FieldName:", field.Name, "allTags:", string(field.Tag))
		// 获取json的tag
		field.Tag.Get("json")
	}
	u.Print("prefix--1")

	fmt.Println("------------------ ")
	// 遍历方法
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println("Method:", t.Method(i).Name)
	}
	fmt.Println("------------------ ")

	// 反射修改字段值
	pv := reflect.ValueOf(&u) // 修改该对象字段，需要传入指针
	puElem := pv.Elem()       // 获取只对象指针的value
	if puElem.IsValid() {
		age2 := 16
		//*pAge = 16
		puElem.FieldByName("Age").Set(reflect.ValueOf(&age2))
	}

	u.Print("prefix--2")

	// 动态调用,注意是Value的MethodByName方法,不是 type的
	printMethod := v.MethodByName("Print")

	if printMethod.IsValid() {
		args := []reflect.Value{reflect.ValueOf("prefix--3")}
		printMethod.Call(args)
	}

	println("结束.....")
}

// ---------------------type---------------------
//  reflect.Zero 函数，可以创建原始类型的对象
func createPrimitiveObjects(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

// 创建数组,不是切片哦
func CreateArray(t reflect.Type, length int) reflect.Value {
	var arrayType reflect.Type
	arrayType = reflect.ArrayOf(length, t)
	return reflect.Zero(arrayType)
}

// 创建chan
func CreateChan(t reflect.Type, buffer int) reflect.Value {
	chanType := reflect.ChanOf(reflect.BothDir, t)
	return reflect.MakeChan(chanType, buffer)
}

// ----------------value的函数----------------

// 从value中取出int ;int32/64同理
func extractInt(v reflect.Value) (int, error) {
	if v.Kind() != reflect.Int {
		return int(0), errors.New("Invalid input")
	}
	var intVal int64
	intVal = v.Int()
	return int(intVal), nil
}

// 获取字符串你
func extractString(v reflect.Value) (string, error) {
	if v.Kind() != reflect.String {
		return "", errors.New("Invalid input")
	}
	return v.String(), nil
}

// Extract Uintptr
func extractUintptr(v reflect.Value) (uintptr, error) {
	if v.Kind() != reflect.Uintptr {
		return uintptr(0), errors.New("Invalid input")
	}
	var ptrVal uintptr
	if v.CanAddr() {
		ptrVal = v.Addr().Pointer()
		return ptrVal, nil
	}
	return uintptr(0), errors.New("can not Extract uintptr")
}

func extractUnsafePointer(v reflect.Value) (unsafe.Pointer, error) {
	if v.Kind() != reflect.UnsafePointer {
		return unsafe.Pointer(uintptr(0)), errors.New("Invalid input")
	}
	var unsafeVal unsafe.Pointer
	unsafeVal = unsafe.Pointer(v.UnsafeAddr())
	return unsafeVal, nil
}

func extractArray(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Array {
		return nil, errors.New("invalid input")
	}
	var array interface{}
	array = v.Interface()
	return array, nil
}

func extractChan(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Chan {
		return nil, errors.New("invalid input")
	}
	var ch interface{}
	ch = v.Interface()
	return ch, nil
}

//----反射创建方法

func CreateFunc(fType reflect.Type, f func(args []reflect.Value) (results []reflect.Value)) (reflect.Value, error) {
	if fType.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("invalid input")
	}

	var ins, outs *[]reflect.Type

	ins = new([]reflect.Type)
	outs = new([]reflect.Type)

	for i := 0; i < fType.NumIn(); i++ {
		*ins = append(*ins, fType.In(i))
	}

	for i := 0; i < fType.NumOut(); i++ {
		*outs = append(*outs, fType.Out(i))
	}
	var variadic bool
	variadic = fType.IsVariadic()
	return AllocateStackFrame(*ins, *outs, variadic, f), nil
}

func AllocateStackFrame(ins []reflect.Type, outs []reflect.Type, variadic bool, f func(args []reflect.Value) (results []reflect.Value)) reflect.Value {
	var funcType reflect.Type
	//FuncOf 方法是用于创建函数的 type signature，
	funcType = reflect.FuncOf(ins, outs, variadic)
	//MakeFunc  方法可以用来给函数分配内存
	return reflect.MakeFunc(funcType, f)
}

// 创建结构体
func CreateStruct(fields []reflect.StructField) reflect.Value {
	var structType reflect.Type
	structType = reflect.StructOf(fields)
	return reflect.Zero(structType)
}

func extractStruct(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.New("invalid input")
	}
	var st interface{}
	st = v.Interface()
	return st, nil
}
