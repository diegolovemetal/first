package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"net"
	"os"
	"time"
)

func main() {
	go startServer("127.0.0.1:8897")
	go startServer("127.0.0.1:8898")
	go startServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<- a
}

func startServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//注册zk节点Q，l链接zk
	conn, err := GetConnect()
	if err != nil {
		fmt.Println("connect zk err:",err)
	}
	defer conn.Close()
	//zk节点注册

	err = RegistServer(conn, port)
	if err != nil {
		fmt.Println("regist node err:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:%s", err)
			continue
		}
		go handleClient(conn, port)
	}
	fmt.Println("aaaaaa")

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnect() (conn *zk.Conn, err error) {
	zklist := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zklist, 10 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error)  {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func handleClient(conn net.Conn, port string){
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime))
	conn.Write([]byte([]byte(port + ":" + daytime)))
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
