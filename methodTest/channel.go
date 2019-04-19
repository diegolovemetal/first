package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int { 12, 23, 45, 61, 34, 45}
	c := make(chan int)
	//channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所
	// 谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任
	// 何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。
	//
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)

	x, y := <-c, <-c
	//FIFO

	fmt.Println(x, y, x+y)


}
