package main

import (
	"bytes"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

type data2 struct {
	num    int               //ok
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", v1 == v2) //prints: v1 == v2: true

	// 使用reflect.DeepEqual 深入比较
	var b1 []byte = nil
	b2 := []byte{}
	fmt.Println("reflect b1 == b2:", reflect.DeepEqual(b1, b2)) //prints: b1 == b2: false
	// bytes.是对于slice nil 和 [] 是相等的
	fmt.Println("bytes b1 == b2:", bytes.Equal(b1, b2)) // true

	vv1 := data2{}
	vv2 := data2{}
	fmt.Println("vv1 == vv2:", reflect.DeepEqual(vv1, vv2))

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true

	// 字符串的比较
	/*
		如果你的byte slice中包含需要验证用户数据的隐私信息（比如，加密哈希、tokens等），
		不要使用reflect.DeepEqual()、bytes.Equal()，或者bytes.Compare()
		因为这些函数将会让你的应用易于被定时攻击 使用'crypto/subtle'包中的函数（即，subtle.ConstantTimeCompare()）
	*/
	zhB := []byte("你好!老外")
	zhS := "你好!老外" // 此处也可能不是utf8类型
	fmt.Println("zhB == zhS", string(zhB) == zhS)
	fmt.Println("zhB strings.EqualFold zhS", strings.EqualFold(string(zhB), zhS))
	fmt.Println("zhB bytes.EqualFold zhS", bytes.EqualFold(zhB, []byte(zhS)))
	fmt.Println("zhB subtle.ConstantTimeCompare zhS", subtle.ConstantTimeCompare(zhB, []byte(zhS)))

	aa1 := []string{"one", "two"}
	aa2 := []interface{}{"one", "two"}
	fmt.Println("aa1 == aa2:", reflect.DeepEqual(aa1, aa2))

	data := map[string]interface{}{
		"code": 200,
		//"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded:", reflect.DeepEqual(data, decoded))
	fmt.Printf("encoded %v \n", reflect.TypeOf(decoded["code"]))

}
