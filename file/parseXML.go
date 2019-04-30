package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type RecurlyServers struct {
	XMLName xml.Name	 `xml:"servers"`
	Version string 		`xml:"version,attr"`
	Svs []server		`xml:"server"`
	Description string	`xml:",innerxml"`
}

type server struct {
	XMLName xml.Name	`xml:"server"`
	ServerName string	`xml:"serverName"`
	ServerIP string		`xml:"serverIP"`
}
func main() {
	file, err := os.Open("src/file/server.xml")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadAll err:", err)
		return
	}
	v := RecurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println(v)

}
