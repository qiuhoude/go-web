package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"192.168.1.151:2377", "192.168.1.151:2378", "192.168.1.151:2379"}
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	showMemberList(cli)
}
func showMemberList(cli *clientv3.Client) {
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	databytes, err := json.MarshalIndent(resp.Members, "", "	")

	fmt.Println("members:", string(databytes))
}

func addMember(cli *clientv3.Client) {
	peerURLs := []string{"http://192.168.1.151:2383"}
	mresp, err := cli.MemberAdd(context.Background(), peerURLs)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("added member.PeerURLs:", mresp.Member.PeerURLs)
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("添加后 members:", resp.Members)
}
