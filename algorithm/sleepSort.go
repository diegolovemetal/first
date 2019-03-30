package main

import (
	"fmt"
	"time"
)

func main() {
	//timer1 := time.NewTimer(time.Second * 2)
	tab := []int{1,233,23,25}

	ch := make(chan  int)
	for _, value := range tab {
		go func(val int) {
			time.Sleep(time.Duration(val) * 10000000)
			fmt.Println(val)
			ch <- val
		}(value)
	}
	for _= range tab {
		<-ch
	}
}
