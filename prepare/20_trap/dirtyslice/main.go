package main

import (
	"bytes"
	"fmt"
)

func main() {

	//testAppend()
	//testRet()
	//testFunParam()

	//testAppend2()

	//testAppend3()

	// 不会发生panic
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println(v)
}

func testAppend3() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5] // len=5-2=3  cap=len(slice)-2=8
	s2 := s1[2:6:7]  // len=6-2=4 cap=7-2=5

	s2 = append(s2, 100)
	// s1[2] = 20  此处修改会s2中首个元素位置的值
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)

	cap := 4
	fmt.Println(uintptr(cap))

}

func testAppend2() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	// 两个文件夹的slice都潜在的引用了同一个原始的路径slice
	//dir1 := path[:sepIndex]

	// 修改:
	// 完整的slice表达式中的额外参数可以控制新的slice的容量。
	// 现在在那个slice后添加元素将会触发一个新的buffer分配，而不是覆盖第二个slice中的数据
	dir1 := path[:sepIndex:sepIndex] //full slice expression

	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1), cap(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2))            //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB (not ok)

	fmt.Println("new path =>", string(path))
}

// 测试函数参数时切片
func testFunParam() {
	raw := []string{"线性代数", "统计", "概率", "微积分"}
	fmt.Printf("before raw:%v %p %v \n", raw, &raw[0], cap(raw))
	//sliceParam(raw)
	sliceParamPrt(&raw)
	fmt.Printf("after raw:%v %p %v \n", raw, &raw[0], cap(raw))

}

func sliceParamPrt(math *[]string) { // 会将 cap len *prt的信息一并传过去
	fmt.Printf("math:%v %p %v \n", *math, &(*math)[0], cap(*math))
	//ns := append(*math, "三角") // 此时创建新底层数组
	//fmt.Printf("ns:%v %p %v \n", ns, &ns[0], cap(ns))

	*math = append(*math, "三角") // 此时创建新底层数组, (注意是指针的重新赋值,会导致调用者的 raw指向也改变)
	fmt.Printf("math:%v %p %v \n", *math, &(*math)[0], cap(*math))
}

func sliceParam(math []string) {
	fmt.Printf("math:%v %p %v \n", math, &math[0], cap(math))
	math = append(math, "三角") // 此时创建新底层数组
	fmt.Printf("math:%v %p %v \n", math, &math[0], cap(math))
}

// 测试返回值对底层数组的改变
func testRet() {
	// 返回回来的切片 都是共享底层数组的
	ret := getSlice()
	ns := ret
	ns2 := append(ret, "三角")
	fmt.Printf("ns:%v %p %d \n", ns, &ns[0], cap(ns))
	fmt.Printf("ns2:%v %p %d \n", ns2, &ns2[0], cap(ns2))

	ret = append(ret, "几何")
	fmt.Printf("ns2:%v %p %d \n", ns2, &ns2[0], cap(ns2))
	fmt.Printf("ret:%v %p %d \n", ret, &ret[0], cap(ret))

}

func getSlice() []string {
	raw := []string{"线性代数", "统计", "概率", "微积分"}
	fmt.Printf("raw: %p\n", &raw[0])
	ret := raw[:3]
	return ret
}

// 测试 append, 对底层数组的改变
func testAppend() {
	raw := []string{"线性代数", "统计", "概率", "微积分"}
	sras := raw[:2]
	ns := append(sras, "机器学习") // 这里底层会创建新的数组

	fmt.Printf("ns:%v %p\n", ns, &ns[0])
	fmt.Printf("raw:%v %p\n", raw, &raw[0])

	raw = append(ns, "均值不等式") // 将 raw指向了 ns的数组, 此时他们两个切片共享一个底层数组
	ns[0] = "golang"

	fmt.Println()
	fmt.Printf("ns:%v %p\n", ns, &ns[0])
	fmt.Printf("raw:%v %p\n", raw, &raw[0])

	fmt.Println()
	ns = append(ns, "复数")
	fmt.Printf("ns:%v %p\n", ns, &ns[0])
	fmt.Printf("raw:%v %p\n", raw, &raw[0])
}
