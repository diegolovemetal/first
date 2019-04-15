package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

// 消费者

type ConsumerT struct {

}

func main() {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}
//处理消息

func (c *ConsumerT) HandleMessage(msg *nsq.Message) error  {
	fmt.Println("receive",msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func InitConsumer(topic string, channel string, add string)  {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second	//设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg)	//新建消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil,0)	//屏蔽系统日志
	c.AddHandler(&ConsumerT{})	//添加消费者接口

	//建立NSQLookupd连接

	if err := c.ConnectToNSQLookupd(add); err != nil {
		panic(err)
	}

}

