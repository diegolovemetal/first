package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 *time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:",err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()

	//设置1秒超时，访问etcd有超时控制
	ctx , cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	_, err = cli.Put(ctx, "logagent/conf", "sample_value")
	//操作完毕，取消etcd
	cancel()
	if err != nil {
		fmt.Println("put failed,err:", err)
		return
	}
	//取值，设置超时一秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx,"logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed ,err:", err)
		return
	}
	for _, val := range resp.Kvs {
		fmt.Printf("%s : %s\n", val.Key, val.Value)
	}

}
