package main

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)


func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	for i:=0; i<100; i++ {
		startClient()

		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	//获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Println("get server host err:",err)
		return
	}

	fmt.Println("connect host" + serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))

	return
}

func getServerHost() (host string, err error) {
	conn, err := GetConnect()
	if err != nil {
		fmt.Println("connect zk err :", err)
		return
	}
	defer conn.Close()
	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Println("get server list error", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty\n")
		return
	}
	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}