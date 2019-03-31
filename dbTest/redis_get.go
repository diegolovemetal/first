package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(" redis connect err:", err)
		return
	}

	defer c.Close()

	_, err = c.Do("MSet", "bcd", 123, "cdf", 234)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "bcd", "cdf"))
	if err != nil {
		fmt.Println("get failed", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}