package demo3_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

const (
	address = "192.168.1.151:6379"
)

//打开redis连接
func open() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()
}

func getSet() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()
	reply, err := conn.Do("set", "language", "golang")
	if err != nil {
		fmt.Println("redis set error:", err)
	} else {
		fmt.Printf("%T replay:%v \n", reply, reply)
	}

	name, err := redis.String(conn.Do("get", "language"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}
}

func mgetmset() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	_, err = conn.Do("MSET", "name", "hanru", "age", 30)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Strings(conn.Do("MGET", "name", "age"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("MGET name: %s \n", res)
		fmt.Println(len(res))
	}
}

func list() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	_, err = conn.Do("LPUSH", "list1", "ele1", "ele2", "ele3", "ele4")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	//res, err := redis.String(conn.Do("LPOP", "list1"))//获取栈顶元素
	//res, err := redis.String(conn.Do("LINDEX", "list1", 3)) //获取指定位置的元素
	res, err := redis.Strings(conn.Do("LRANGE", "list1", 0, -1)) //获取指定下标范围的元素
	if err != nil {
		fmt.Println("redis POP error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %s \n", res)
	}
}

func hash() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	_, err = conn.Do("HSET", "user", "name", "hanru", "age", 30)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Int64(conn.Do("HGET", "user", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %d \n", res)
	}
}

func pipelining() {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	// 使用 pipeline
	// Send：发送命令至缓冲区
	conn.Send("HSET", "user", "name", "hanru", "age", "30")
	conn.Send("HSET", "user", "sex", "female")
	conn.Send("HGET", "user", "age")
	// 清空缓冲区，将命令一次性发送至服务器
	conn.Flush()

	// Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)

}
