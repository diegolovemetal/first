package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file , err := os.Open("src/file1.txt")
	if err != nil {
		fmt.Printf("open err:%v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	data1, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ioutil readALL err:%v\n", err)
	}
	fmt.Printf("ioutil read all success.\n内容:\n %s \n",string(data1))
	data2, err := ioutil.ReadFile("src/file1.txt")
	if err != nil {
		fmt.Printf("ioutil read file err:%v\n", err)
	}
	fmt.Printf("ioutil read file success.\n内容:\n%s\n", string(data2))
}
