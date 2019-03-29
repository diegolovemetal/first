package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str := "this is ioutil.writeFile() test\n"
	err := ioutil.WriteFile("./ioutilWriteFile.txt", []byte(str), os.ModePerm)

	if err != nil {
		fmt.Printf("ioutil.WriteFile() write./ioutilWriteFile.txt err:%v\n", err)
	}
	fmt.Println("ioutil.WriteFile() write ./ioutilWriteFile.txt success")
}
