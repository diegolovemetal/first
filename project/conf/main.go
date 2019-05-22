package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main()  {
	config, err := config.NewConfig("ini", "src/project/conf/test.conf")
	if err != nil {
		fmt.Println("new config failed, err: ", err)
		return
	}

	appname := config.String("appname")

	fmt.Println("get app name:", appname)

	httpport, err := config.Int("httpport")

	if err != nil {
		fmt.Printf("get httpport err: %v\n", err)
	}
	fmt.Println("get httpport :", httpport)

	mysqlport, err := config.Int("mysqlport")

	if err != nil {
		fmt.Printf("get mysqlport err: %v\n", err)
	}
	fmt.Println("get mysqlport :", mysqlport)

	pi, err := config.Float("PI")
	if err != nil {
		fmt.Println("get PI err:", err)
	}
	fmt.Println("get PI:", pi)

	AutoRender, err := config.Bool("autorender")
	if err != nil {
		fmt.Println("get autorender err:", err)
	}
	fmt.Println("get autorender:", AutoRender)

	key1 := config.String("demo::key1")
	key2 := config.String("demo::key2")
	fmt.Println("key1 and key2 are:", key1, key2)

}
