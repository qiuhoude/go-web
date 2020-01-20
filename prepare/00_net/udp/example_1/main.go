package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/prometheus/common/log"
	"net"
)

func server() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	fmt.Printf("Local: <%s> \n", conn.LocalAddr().String())

	data := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("server : rsv <%s> %s\n", remoteAddr, data[:n])

		_, err = conn.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}

func client() {
	sip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 8888}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(`hello`))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<%s>\n", conn.RemoteAddr())
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	fmt.Printf("client :read %s from <%s>\n", data[:n], conn.RemoteAddr())

}

func main() {

	go server()

	client()

	console.ReadLine()
}
