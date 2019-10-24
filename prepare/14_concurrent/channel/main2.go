package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
		fmt.Println("结束")
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	close(ch)
	time.Sleep(time.Second)
}
