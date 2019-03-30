package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	//return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano()).Int63(1000000)))
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

}
func main() {
	fmt.Println(CreateCaptcha())
}

