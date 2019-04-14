package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/*
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.iNewSyncProducer
	创建消息sarama.ProducerMessage
	发送消息client.SendMessage
*/
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test")

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}

	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed:,err:", err)
		return
	}
	fmt.Printf("pid %v: offset:%v\n", pid, offset)
}
