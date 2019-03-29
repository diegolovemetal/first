package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i:=0; i<50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&sum, 1)
			//sum += 1
		}()
		wg.Wait()
		fmt.Println(sum)
	}
	//for i:=0; i<50; i++ {
	//	sum += 1
	//}
	//fmt.Println(sum)
}
