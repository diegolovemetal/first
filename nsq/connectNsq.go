package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func main() {
	//定义NSQ生产者
	var producer *nsq.Producer
	//初始化生产者
	// producer, err := nsq.NewProducer("地址:端口", nsq.*Config )
	producer, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	err = producer.Ping()
	if err != nil {
		//关闭生产者

		producer.Stop()
		producer = nil
	}
	fmt.Println("ping nsq success")

}
