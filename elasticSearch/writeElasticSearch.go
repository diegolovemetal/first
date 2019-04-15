package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User string
	Message string
}
func main() {
	//创建ES客户端
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		fmt.Println("connect ES error :",err)
		return
	}
	//创建成功
	fmt.Println("connect ES success")

	//写入数据
	for i:=0; i<20; i++ {
		tweet := Tweet{User:"diego", Message:" love Sasa"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do()
	}
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("insert Success")
}
