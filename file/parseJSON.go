package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"Servers":[{"ServerName": "ShangHai_VPN", "ServerIP":"127.0.0.1"}, {"ServerName": "BeiJing_VPN", "ServerIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}