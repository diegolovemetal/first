package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init(){
	pool = &redis.Pool{ //实例化连接池
		MaxIdle:16, //最初连接数量
		MaxActive:0, //最大连接数量，0表示自动
		IdleTimeout:300, //连接关闭时间300s，
		Dial: func() (conn redis.Conn, e error) {	//要连接的数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	c := pool.Get()	//从连接池，取一个连接
	defer c.Close() //函数运行结束，将连接放回连接池

	_, err := c.Do("Set","abc",100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,err:" ,err)
		return
	}
	fmt.Println(r)
	pool.Close()

}