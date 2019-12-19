package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func pathTest() {
	//返回路径的最后一个元素
	fmt.Println(path.Base("./github.com/mojocn/c"))
	//如果路径为空字符串,返回.
	fmt.Println(path.Base(""))
	//如果路径只有斜线,返回/
	fmt.Println(path.Base("///"))

	//返回等价的最短路径
	//1.用一个斜线替换多个斜线
	//2.清除当前路径.
	//3.清除内部的..和他前面的元素
	//4.以/..开头的,变成/
	fmt.Println(path.Clean("./github.com/mojocn/../"))

	//返回路径最后一个元素的目录
	//路径为空则返回.
	fmt.Println(path.Dir("./github.com/mojocn/c"))

	//返回路径中的扩展名
	//如果没有点,返回空
	fmt.Println(path.Ext("./github.com/mojocn/c/d.jpg"))

	//判断路径是不是绝对路径
	fmt.Println(path.IsAbs("./github.com/mojocn/c"))
	fmt.Println(path.IsAbs("/github.com/mojocn/c"))

	//连接路径,返回已经clean过的路径
	fmt.Println(path.Join("./a", "b/c", "../d/"))

	//匹配文件名,完全匹配则返回true
	fmt.Println(path.Match("*", "a"))
	fmt.Println(path.Match("*", "a/b/c"))
	fmt.Println(path.Match("\\b", "b"))

	//分割路径中的目录与文件
	fmt.Println(path.Split("./github.com/mojocn/c/d.jpg"))
}

func filePathTest() {
	//返回所给路径的绝对路径
	path, _ := filepath.Abs("./1.txt")
	fmt.Println(path)

	//返回路径最后一个元素
	fmt.Println(filepath.Base("./1.txt"))
	//如果路径为空字符串,返回.
	fmt.Println(filepath.Base(""))
	//如果路径只有斜线,返回/
	fmt.Println(filepath.Base("///"))

	//返回等价的最短路径
	//1.用一个斜线替换多个斜线
	//2.清除当前路径.
	//3.清除内部的..和他前面的元素
	//4.以/..开头的,变成/
	fmt.Println(filepath.Clean("C:/github.com/mojocn/../c"))
	fmt.Println(filepath.Clean("./1.txt"))

	//返回路径最后一个元素的目录
	//路径为空则返回.
	fmt.Println(filepath.Dir("./github.com/mojocn/c"))
	fmt.Println(filepath.Dir("C:/github.com/mojocn/c"))

	//返回链接文件的实际路径
	path2, _ := filepath.EvalSymlinks("1.lnk")
	fmt.Println(path2)

	//返回路径中的扩展名
	//如果没有点,返回空
	fmt.Println(filepath.Ext("./github.com/mojocn/c/d.jpg"))

	//将路径中的/替换为路径分隔符
	fmt.Println(filepath.FromSlash("./github.com/mojocn/c"))

	//返回所有匹配的文件
	match, _ := filepath.Glob("./*.go")
	fmt.Println(match)

	//判断路径是不是绝对路径
	fmt.Println(filepath.IsAbs("./github.com/mojocn/c"))
	fmt.Println(filepath.IsAbs("C:/github.com/mojocn/c"))

	//连接路径,返回已经clean过的路径
	fmt.Println(filepath.Join("C:/a", "/b", "/c"))

	//匹配文件名,完全匹配则返回true
	fmt.Println(filepath.Match("*", "a"))
	fmt.Println(filepath.Match("*", "C:/github.com/mojocn/c"))
	fmt.Println(filepath.Match("\\b", "b"))

	//返回以basepath为基准的相对路径
	path3, _ := filepath.Rel("C:/github.com/mojocn", "C:/github.com/mojocn/c/d/../e")
	fmt.Println(path3)

	//将路径使用路径列表分隔符分开,见os.PathListSeparator
	//linux下默认为:,windows下为
	fmt.Println(filepath.SplitList("C:/windows  C:/windows/system"))

	//分割路径中的目录与文件
	dir, file := filepath.Split("C:/github.com/mojocn/c/d.jpg")
	fmt.Println(dir, file)

	//将路径分隔符使用/替换
	fmt.Println(filepath.ToSlash("C:/github.com/mojocn"))

	//返回分区名
	fmt.Println(filepath.VolumeName("C:/github.com/mojocn/c"))

}

/*
遍历根目录(root)下的文件树,为树中的每个文件或目录(包括根目录)调用walkFn.
所有在访问文件和目录时出现的错误都由walkFn过滤. 遍历按词法顺序进行,
这使得输出是确定的,但对于非常大的目录来说,遍历可能是低效的. filepath.Walk()不会跟进符号链接.
*/
func walkTest() {
	layout := "2006-01-02 15:04:05"

	//遍历指定目录下所有文件
	_ = filepath.Walk(".", func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // can't walk here,
			return nil       // but continue walking elsewhere
		}
		if fi.IsDir() {
			return nil // not a file.  ignore.
		}
		// 过滤输出内容
		matched, err := filepath.Match("*.txt", fi.Name())
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			// fmt.Println(fp)
			fmt.Printf("Name: %s, ModifyTime: %s, Size: %v\n", fp, fi.ModTime().Format(layout), fi.Size())
		}
		return nil
	})
}

/*
filepath.Walk()会自动遍历子目录,但有些时候我们不希望这样,如果只想看当前目录,
或手动指定某几级目录中的文件,这个时候,可以使用 ioutil.ReadDir 进行替代.
*/
func readDirTest(path string, level int) {
	readFile(path, 0, level)
}
func readFile(p string, curLv, level int) {
	if curLv > level {
		return
	}
	fi, err := os.Stat(p)
	if err != nil {
		return
	}
	var sb strings.Builder
	for i := 0; i < curLv; i++ {
		sb.WriteString(" ")
	}
	fmt.Printf("%v%v\n", sb.String(), fi.Name())
	if fi.IsDir() {
		files, err := ioutil.ReadDir(p)
		if err != nil {
			log.Println(err)
			return
		}
		for _, fi := range files {
			readFile(path.Join(p, fi.Name()), curLv+1, level)
		}
	}

}

func main() {
	//pathTest()
	//filePathTest()
	//walkTest()
	//readDirTest(".", 2)

}
