package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	do(1 * time.Second)
	fmt.Println("exit")
}

func do(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("超时了")
			return
		default:
			//time.Sleep(3 * time.Second)
			//fmt.Println("执行完毕")
		}
	}

}
