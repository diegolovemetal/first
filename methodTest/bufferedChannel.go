package main

import "fmt"

func main() {
	//c := make(chan int, 2)
	//c <- 2
	//c <- 3
	//
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	c := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <-0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	//for i:=0; i<n; i++ {
	//	c <- x
	//	x, y = y, x+y
	//}
	//close(c)
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

