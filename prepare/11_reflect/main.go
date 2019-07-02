package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/11_reflect/rtype"
	"reflect"
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
