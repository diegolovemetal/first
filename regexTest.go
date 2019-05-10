package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	//ip1 := "12312.sdfas.sd12"
	//ip2 := "192.168.1.1"
	//
	//fmt.Println(IsIP(ip1))
	//fmt.Println(IsIP(ip2))
	//判断是否为IP地址

	if len(os.Args) == 1 {
		fmt.Println("Usage: regex[string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("是个数字")
	} else {
		fmt.Println("不是数字")
	}
}
