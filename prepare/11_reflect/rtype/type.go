package rtype

import "fmt"

type User struct {
	Name string      `json:"name" bson:"b_name"`
	Age  *int        `json:"age"`
	Love interface{} `json:"love"`
	lv   int         `json:"lv"`
}

func (u User) GetLv() int {
	return u.lv
}

func (u *User) SetLv(lv int) {
	u.lv = lv
}

func (u User) Print(prfix string) {
	if u.Age == nil {
		// 对nil取 * 是会panic的
		fmt.Printf("%s:Name is %s,Age is ? ageAddr=%v Love=%v lv=%d \n", prfix, u.Name, u.Age, u.Love, u.lv)
	} else {
		fmt.Printf("%s:Name is %s,Age is %v ageAddr=%v Love=%v lv=%d \n", prfix, u.Name, *u.Age, u.Age, u.Love, u.lv)
	}
}

type User1 struct {
	B   byte
	I32 int32
	I64 int64
}
