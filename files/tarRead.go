package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	fileName := "./file.tar.gz"
	file ,err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file ./file.tar.gz err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	read := tar.NewReader(file)
	hdr, err := read.Next()
	var getByte = make([]byte, hdr.Size)
	_, err = read.Read(getByte)
	if err != nil {
		fmt.Printf("read err : %v\n", err)

	}
	fmt.Println(string(getByte))
}
