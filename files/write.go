package main

import (
	"fmt"
	"os"
)

func main(){
	file1, err := os.Open("./file1.txt")
	if err != nil {
		fmt.Printf("open file1 error:%v\n", err)
	}
	if file1 != nil {
		defer func(file *os.File) { file.Close() }(file1)
	}

	file2, err := os.OpenFile("./file2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("oepn file2 err:%v\n", err)
	}
	if file2 != nil {
		defer func(file *os.File) { file.Close() }(file2)
	}

	b1 := []byte("hello world,write test! \n")
	off, err := file2.Write(b1)
	if err != nil {
		fmt.Printf("write file2 err:%v\n", err)
	}
	fmt.Printf("file2 writes success, off is %v\n", off)

	b2 := []byte("hello golang! writeAt test!\n")
	off, err = file2.WriteAt(b2, int64(off))
	if err != nil {
		fmt.Printf("file2 writeAt err:%v\n",err)
	}
	fmt.Printf("file2 writeAt success, off is %v\n", off)

	str := "this is write string test!\n"
	off, err = file2.WriteString(str)
	if err != nil {
		fmt.Printf("file2 write string err :%v\n", err)
	}
	fmt.Printf("file2 write string test success,off is %v\n", off)
}
