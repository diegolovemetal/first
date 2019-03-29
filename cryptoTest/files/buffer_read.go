package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("src/file1.txt")
	if err != nil {
		fmt.Printf("os open err:%v\n", err)
	}
	if file != nil {
		defer func(file *os.File) {file.Close()}(file)
	}

	read1 := bufio.NewReader(file)

	b1 := make([]byte, 102)
	readByte1, err := read1.Read(b1)
	if err != nil {
		fmt.Printf("read err:%d\n", err)
	}
	fmt.Printf("read success, 读取 %d 字节\n 读取的内容: \n %s\n", readByte1, string(b1))

	var line []byte
	for {
		data, prefix, err := read1.ReadLine()
		if err == io.EOF {
			break
		}

		line = append(line, data...)
		if prefix {
			fmt.Println("没有读取完")
		}

	}
	fmt.Println(string(line))
}
