package main

import (
	"bufio"
	"fmt"
	"os"
)

type closeer struct {
	a int
	b string
}

func main() {

	ch := make(chan closeer)

	for i := 0; i < 10; i++ {
		id := i
		go func() {
			defer func() {
				fmt.Printf("end goroutineId:%d\n", id)
			}()
			select {
			case data := <-ch:
				fmt.Printf("recve goroutineId:%d, data:%v \n", id, data)
				return
			}
		}()
	}
	//ch <- struct{}{}

	fmt.Println("send")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	close(ch)
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
