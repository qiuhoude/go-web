package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	address = "192.168.1.151:6379"
)

func main() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("Subs connect success ...")
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
}
