package main

import (
	"fmt"
	"net"
)

func main() {
	//ipstr := "192.168.10.105"
	ipstr := "2001:0db8:85a3:08d3:1319:8a2e:0370:7344"
	addr := net.ParseIP(ipstr)

	if addr != nil {
		fmt.Println("ip is", addr.String())
	} else {
		fmt.Println("Invalid address")
	}
}
