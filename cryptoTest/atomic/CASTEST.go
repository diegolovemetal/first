package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i:= uint32(0); i<100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.CompareAndSwapUint32(&sum, 100, sum+1)
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
