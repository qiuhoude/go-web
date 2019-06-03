package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func register(w http.ResponseWriter, r *http.Request) {

	//r.FormValue("username") 这个可以不需要提前调用 ParseForm()
	// 只会返回同名参数中的第一个，若参数不存在则返回空字符串。

	r.ParseForm()
	//1.验证必填字段
	username := r.Form.Get("username")
	if len(username) == 0 {
		fmt.Println("用户名不能为空！")
		fmt.Fprintf(w, "用户名不能为空！") //这个写入到w的是输出到客户端的
	}
	//2.验证数字
	age, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		//数字转化出错了，那么可能就是不是数字
		fmt.Println("您输入的不是数字！")
		fmt.Fprintf(w, "您输入的不是数字！") //这个写入到w的是输出到客户端的
	}
	//接下来就可以判断这个数字的大小范围了
	if age > 100 || age < 0 {
		//太大了或太小了
		fmt.Println("您输入的年龄太大了或太小了，请输入0-100之间的整数！")
		fmt.Fprintf(w, "您输入的年龄太大了或太小了，请输入0-100之间的整数！") //这个写入到w的是输出到客户端的
	}
	//或者正则表达式
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		fmt.Println("验证有误，您输入的年龄太大了或太小了！")
		fmt.Fprintf(w, "验证有误，您输入的年龄太大了或太小了！")
	}

	//3.验证中文
	if m, _ := regexp.MatchString(`^[\x{4e00}-\x{9fa5}]+$`, r.Form.Get("zhname")); !m {
		fmt.Println("验证有误，请输入中文！")
		fmt.Fprintf(w, "验证有误，请输入中文！")
	}

	//4. 验证英文
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("enname")); !m {
		fmt.Println("验证有误，请输入英文！")
		fmt.Fprintf(w, "验证有误，请输入英文！")
	}

	//5. 邮箱
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("请输入正确邮箱地址")
		fmt.Fprintf(w, "验证有误，请输入正确邮箱地址！")
	}

	//6. 验证手机号
	if m, _ := regexp.MatchString(`^(1[3|5|6|7|8][0-9]\d{8})$`, r.Form.Get("mobile")); !m {
		fmt.Println("请输入正确手机号码")
		fmt.Fprintf(w, "验证有误，请输入正确手机号码！")
	}

	//7. 下拉菜单
	xueli := r.Form.Get("xueli")
	res1 := checkSelect(xueli)
	if !res1 {
		fmt.Println("请选择正确的下拉列表！")
		fmt.Fprintf(w, "请选择正确的下拉列表！")
	}

	// 8. 单选按钮
	sex := r.Form.Get("sex")
	res2 := checkSex(sex)
	if !res2 {
		fmt.Println("请选择正确的性别！")
		fmt.Fprintf(w, "请选择正确的性别！")
	}

	// 9. 复选框
	hobby := r.Form["hobby"]
	res3 := checkHobby(hobby)
	if !res3 {
		fmt.Println("请选择正确的爱好！")
		fmt.Fprintf(w, "请选择正确的爱好！")
	}
	// 10 身份证号
	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		fmt.Println("请选择正确的身份证号！")
		fmt.Fprintf(w, "请选择正确的身份证号！")
	}

	//fmt.Println("验证成功！")
	//fmt.Fprintf(w, "验证成功！")
}

func checkHobby(hobby []string) bool {
	slice := []string{"game", "girl", "money", "power"}
	hobby2 := Slice_diff(hobby, slice)
	if hobby2 == nil {
		return true
	}
	return false
}

func Slice_diff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

/**
判断是一个切片中是否包含指定的数值
*/
func InSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

/**
验证单选按钮
*/
func checkSex(sex string) bool {
	slice := []string{"male", "female", "other"}
	for _, v := range slice {
		if v == sex {
			return true
		}
	}
	return false
}

/**
验证下拉列表
*/
func checkSelect(xueli string) bool {
	slice := []string{"xiaoxue", "chuzhong", "gaozhong", "dazhuan", "benke", "shuoshi", "boshi", "lieshi"}
	for _, v := range slice {

		if v == xueli {
			return true
		}
	}
	return false
}
func main() {
	http.HandleFunc("/register", register)   //设置访问的路由
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
验证数字：^[0-9]*$
验证n位的数字：^\d{n}$
验证至少n位数字：^\d{n,}$
验证m-n位的数字：^\d{m,n}$
验证零和非零开头的数字：^(0|[1-9][0-9]*)$
验证有两位小数的正实数：^[0-9]+(.[0-9]{2})?$
验证有1-3位小数的正实数：^[0-9]+(.[0-9]{1,3})?$
验证非零的正整数：^\+?[1-9][0-9]*$
验证非零的负整数：^\-[1-9][0-9]*$
验证非负整数（正整数 + 0） ^\d+$
验证非正整数（负整数 + 0） ^((-\d+)|(0+))$
验证长度为3的字符：^.{3}$
验证由26个英文字母组成的字符串：^[A-Za-z]+$
验证由26个大写英文字母组成的字符串：^[A-Z]+$
验证由26个小写英文字母组成的字符串：^[a-z]+$
验证由数字和26个英文字母组成的字符串：^[A-Za-z0-9]+$
验证由数字、26个英文字母或者下划线组成的字符串：^\w+$
验证用户密码:^[a-zA-Z]\w{5,17}$ 正确格式为：以字母开头，长度在6-18之间，只能包含字符、数字和下划线。
验证是否含有 ^%&',;=?$\" 等字符：[^%&',;=?$\x22]+
验证汉字：^[\u4e00-\u9fa5],{0,}$
验证Email地址：^\w+[-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
验证InternetURL：^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$ ；^[a-zA-z]+://(w+(-w+)*)(.(w+(-w+)*))*(?S*)?$
验证电话号码：^(\(\d{3,4}\)|\d{3,4}-)?\d{7,8}$：--正确格式为：XXXX-XXXXXXX，XXXX-XXXXXXXX，XXX-XXXXXXX，XXX-XXXXXXXX，XXXXXXX，XXXXXXXX。
验证身份证号（15位或18位数字）：^\d{15}|\d{}18$
验证一年的12个月：^(0?[1-9]|1[0-2])$ 正确格式为：“01”-“09”和“1”“12”
验证一个月的31天：^((0?[1-9])|((1|2)[0-9])|30|31)$ 正确格式为：01、09和1、31。
整数：^-?\d+$
非负浮点数（正浮点数 + 0）：^\d+(\.\d+)?$
正浮点数 ^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$
非正浮点数（负浮点数 + 0） ^((-\d+(\.\d+)?)|(0+(\.0+)?))$
负浮点数 ^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*)))$
浮点数 ^(-?\d+)(\.\d+)?$
*/
