package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func RedisInit(addr string) (*redis.Client,error) {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     	addr,
		Password: 	"",
		DB:       	0,
		PoolSize:	10,
	})
	_,err := redisdb.Ping().Result()
	if err!=nil {
		return nil,err
	}else{
		return redisdb,nil
	}
}

func ListTest(redisdb *redis.Client) {
	result, err := redisdb.RPush("list1", "a", "b", "c").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func DecrTest(redisdb *redis.Client) {
	redisdb.Set("key01", "10", 0)
	result, err := redisdb.Decr("key01").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func DecrByTest(redisdb *redis.Client) {
	redisdb.Set("key01", "10", 0)
	result, err := redisdb.DecrBy("key01", 3).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
func AppendTest(redisdb *redis.Client) {
	redisdb.Set("key01", "abc", 0)
	result, err := redisdb.Append("key01", "def").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}


func main()  {
	redisdb,err :=RedisInit("localhost:6379")
	if err != nil {
		fmt.Println("init err:",err)
		return
	}
	ListTest(redisdb)
	DecrTest(redisdb)
	DecrByTest(redisdb)
	AppendTest(redisdb)
	}