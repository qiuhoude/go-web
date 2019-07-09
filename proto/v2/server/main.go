package main

import (
	"errors"
	"runtime"
	"time"
)
import "fmt"

func main() {
	ticker := time.NewTicker(1<<63 - 1)
	defer ticker.Stop()
	done := make(chan bool)
	TraceCode(errors.New("二日"))
	time.AfterFunc(10*time.Second, func() {
		done <- true
	})
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	done <- true
	//}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}

}

func TraceCode(code ...interface{}) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	data := ""
	for _, v := range code {
		data += fmt.Sprintf("%v", v)
	}

	data += string(buf[:n])
	fmt.Printf("==> %s\n", data)
}
