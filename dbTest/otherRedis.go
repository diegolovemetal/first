package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
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
func SetTest(redisdb *redis.Client) {
	err:=redisdb.Set("key1","hello",0 * time.Second).Err()
	if err!=nil {
		fmt.Println(err)
		return
	}

	err=redisdb.Set("key2","world",0 * time.Second).Err()
	if err!=nil {
		fmt.Println(err)
		return
	}
}
func GetTest(redisdb *redis.Client) {
	result,err:=redisdb.Get("key1").Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
func GetSetTest(redisdb *redis.Client) {
	result, err := redisdb.GetSet("key1", "我改了，好了").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func MgetTest(redisdb *redis.Client) {
	result, err := redisdb.MGet("key1", "key2","key3","key4").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
	}

func SetnxTest(redisdb *redis.Client)  {
	result, err := redisdb.SetNX("key1", "diego", 0).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
func MsetTest(redisdb *redis.Client) {
	result,err:=redisdb.MSet("key4","h4","key2","h2","key3","h3").Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func MsetnxTest(redisdb *redis.Client) {
	result, err := redisdb.MSetNX("key5", "value5", "key6", "value6").Result()
	/* 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。

即使只有一个给定 key 已存在， MSETNX 也会拒绝执行所有给定 key 的设置操作。

MSETNX 是原子性的，因此它可以用作设置多个不同 key 表示不同字段(field)的唯一性逻辑对象(unique logic object)，所有字段要么全被设置，要么全不被设置。
	所以第一次操作为true，第二次操作为false
	*/
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
//名称为key的string增1操作。
func IncrTest(redisdb *redis.Client) {
	redisdb.Set("key1","1",0)
	result, err := redisdb.Incr("key1").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
//名称为key的string增加integer。
func IncrByTest(redisdb *redis.Client) {
	//设置key1-1
	redisdb.Set("key1","1",0)
	result,err:=redisdb.IncrBy("key1",10).Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
//append(key, value)：名称为key的string的值附加value。
func AppendTest(redisdb *redis.Client) {
	redisdb.Set("key1","abc",0)
	result,err:=redisdb.Append("key1","defg").Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func main()  {
	redisdb,err:=RedisInit("localhost:6379")
	if err != nil {
		fmt.Println("init err:",err)
		return
	}
	SetTest(redisdb)
	GetTest(redisdb)
	GetSetTest(redisdb)
	SetnxTest(redisdb)	//因为key1存在所以返回false
	MsetTest(redisdb)
	MgetTest(redisdb)
	MsetnxTest(redisdb)
	IncrTest(redisdb)
	IncrByTest(redisdb)
	AppendTest(redisdb)  //返回7，append后字符串"abcdef"长度为6
}