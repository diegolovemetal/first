package main

import (
	"fmt"
	"os"
)

func main() {
	file, err  := os.Open("src/file1.txt")
	if err != nil {
		fmt.Printf("Open file1.txt err:%v\n ", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	b1 := make([]byte, 102)
	Space1, err := file.Read(b1)
	if err != nil {
		fmt.Printf("file read err:%v\n", err)
	}
	fmt.Printf("file read success ,读取了 %d 字节\n", Space1)
	fmt.Printf("读取内容 :%s \n",string(b1))

	b2 := make([]byte, 205)
	Space2, err := file.ReadAt(b2, int64(Space1))
	if err != nil {
		fmt.Printf("readAt err:%v\n",err)
	}
	fmt.Printf("file readAt success,读取了%d字节。\n",Space2)
	fmt.Printf("读取内容:%s\n", string(b2))

}
