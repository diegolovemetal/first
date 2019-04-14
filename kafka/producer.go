package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err :",err)
		return
	}
	defer client.Close()

	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a test")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,err:", err)
			return
		}

		fmt.Printf("pid: %v offset %v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)

	}
}
