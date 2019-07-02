// 求下面f 的返回值
// defer是在return之前执行的, return xxx这一条语句并不是一条原子指令!

// 将返回值进行拆分
/*
返回值 = xxx
调用defer函数
空的return
*/

package main

import "fmt"

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

/*
func f1() (result int) {
 	result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	defer func() {//defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return //空的return
}
*/

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

/*
func f2() (r int) {
	t := 5
	r = t 	//defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
	defer func() {
		t = t + 5
	}()
	return			// 空指令
}
*/

func f3() (r int) {
	defer func(r int) { //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return 1
}

/*

func f3() (r int) {
	r = 1  //给返回值赋值
	defer func(r int) {
		r = r + 5
	}(r)
	return 	//空的return
}

*/
