package main

import (
	"fmt"
	"sync/atomic"
)

var old_val uint32 = 10
var new_val uint32 = 100

func main() {
	//atomic.StoreUint32(&val, 100)
	atomic.SwapUint32(&old_val, new_val)
	fmt.Println(old_val, new_valx)
}
