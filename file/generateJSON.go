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
	s.Servers = append(s.Servers, Server{ServerName:"BeiJing", ServerIP:"127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName:"ShangHai", ServerIP:"127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

//我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？
// JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，
// 所以必须通过struct tag定义来实现：
//
//type Server struct {
//	ServerName string `json:"serverName"`
//	ServerIP   string `json:"serverIP"`
//}
//
//type Serverslice struct {
//	Servers []Server `json:"servers"`
//}
