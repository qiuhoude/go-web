package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool redis.Pool

const (
	address = "192.168.1.151:6379"
)

func init() {
	pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
	}
}

func main() {
	conn := pool.Get()
	res, err := conn.Do("HSET", "user", "name", "hanru")
	fmt.Printf("res:%v, error:%v \n", res, err)
	res1, err := redis.String(conn.Do("HGET", "user", "name"))
	fmt.Printf("res1:%s, error:%v \n", res1, err)

}
