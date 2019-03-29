//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var (
//	m = make(map[int]uint64)
//	lock sync.Mutex
//)
//
//type task struct {
//	n int
//}
//
//func calc(t *task){
//	var sum uint64
//	sum = 1
//	for i:=1; i<t.n; i++ {
//		sum *= uint64(i)
//	}
//	fmt.Println(t.n, sum)
//	lock.Lock()
//	m[t.n] = sum
//	lock.Unlock()
//	}
//
//func main() {
//	for i:=0; i<16; i++ {
//		t := &task{n : i}
//		go calc(t)
//	}
//
//	time.Sleep(10 * time.Second)
//	lock.Lock()
//	for k,v := range m {
//		fmt.Printf("%d! = %v\n", k-1, v)
//	}
//	lock.Unlock()
//}
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("lock the lock.(G0)")
	mutex.Lock()

	fmt.Println("the lock is locked.(G0)")
	for i:=1; i<4; i++ {
		go func(i int) {
			fmt.Printf("lock the lock.(G%d)\n", i)
			mutex.Lock()
			fmt.Printf("the lock is locked.(G%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("unlock the lock.(G0)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (G0)")
	time.Sleep(time.Second)
}