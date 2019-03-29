package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var rwlock sync.RWMutex
var lock sync.Mutex

//func testRWLock() {
//	a := make(map[int]int, 5)
//	a[8] = 10
//	a[1] = 10
//	a[2] = 10
//	a[3] = 10
//	a[4] = 10
//	for i:=0; i<2; i++ {
//		go func(b map[int]int) {
//			rwlock.Lock()
//			b[8] = rand.Intn(100)
//			rwlock.Unlock()
//		}(a)
//	}
//
//	for i:=0; i<10; i++ {
//		go func(b map[int]int) {
//			rwlock.RLock()
//			fmt.Println(a)
//			rwlock.RUnlock()
//		}(a)
//	}
//	time.Sleep(2 * time.Second )
//}
func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for{
				rwlock.RLock()
			 //读锁
			 	time.Sleep(time.Millisecond)
				rwlock.RUnlock()
				atomic.AddInt32(&count, 1)

			}
		}(a)
	}
	time.Sleep(time.Second * 20)
	fmt.Println(atomic.LoadInt32(&count))

}

func main() {
	testRWLock()
}


