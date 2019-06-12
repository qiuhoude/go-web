package json

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
)

func TestParseJson(t *testing.T) {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func TestParseJsonMap(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}

	err := json.Unmarshal(b, &f)
	if err != nil {
		t.Fatal(err)
	}

	m, ok := f.(map[string]interface{})
	if !ok {
		t.Fatal("转换失败")
	}
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func TestSimpleJosnDemo(t *testing.T) {
	js, err := simplejson.NewJson([]byte(`{
		"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
		}
		}`))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("%v\n", arr)
	fmt.Printf("%d\n", i)
	fmt.Printf("%s\n", ms)
}

func TestCreateJson(t *testing.T) {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
