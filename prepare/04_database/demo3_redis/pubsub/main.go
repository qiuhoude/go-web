package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	address     = "192.168.1.151:6379"
	channelName = "testChannel1"
)

// 订阅
func Subs() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("Subs connect success ...")
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channelName) //订阅频道

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("Subscription -> %s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(msg string) {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("Push connect success ...")
	defer conn.Close()

	if reply, err := conn.Do("PUBLISH", channelName, msg); err != nil {
		fmt.Println("pub err: ", err)
	} else {
		fmt.Printf("reply:%v \n", reply)
	}

}

func main() {
	go Subs()
	go Push("lala lu lala")
	time.Sleep(time.Second * 3)
}
