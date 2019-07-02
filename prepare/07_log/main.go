package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("【debug】")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Llongfile) // 其实使用按位运行来选
}

func main() {
	log.Println("haha") // 会调用std.Output ,向控制台输出

	//b := []byte(nil)
	aa := fmt.Sprintf(`%v`, nil)
	fmt.Println(aa)

	// Println writes to the standard logger.
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	log.Panicln("panic message")

}
