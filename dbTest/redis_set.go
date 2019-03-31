package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed:",err)
		return
	}

	defer c.Close()
	_, err = c.Do("Set", "abc", 123)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)

	_, err = c.Do("expire", "abc",10)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = c.Do("lpush", "book_list", "math", "english", "chemistry")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := redis.String(c.Do("lpop","book_list"))
	if err != nil {
		fmt.Println("get list failed,", err)
		return
	}

	fmt.Println(result)

	_,err = c.Do("Hset", "books", "abc", 1234)
	if err != nil {
		fmt.Println(err)
		return
	}

	res ,err := redis.Int(c.Do("Hget", "books", "abc"))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
