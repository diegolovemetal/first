package main

import (
	"encoding/xml"
	"fmt"

	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"ShangHai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"BeiJing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Println("error", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
