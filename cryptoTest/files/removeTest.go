package main

import (
	"fmt"
	"os"
)

func main() {
	//err := os.Remove("./file2.txt")
	//if err != nil{
	//	fmt.Printf("remove ./file2.txt err:%v\n", err)
	//}

	err := os.RemoveAll("./file1.txt")
	if err != nil{
		fmt.Printf("remove ./file1.txt err:%v\n", err)
	}
}