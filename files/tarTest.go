package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./file.tar.gz"
	inserByte := []byte("this is test tar write")
	file ,err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file ./file.tar.gz err:%v\n", err)
	}
	if file != nil{
		defer func(file *os.File) { file.Close()}(file)
	}

	write := tar.NewWriter(file)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("os stat err :%v\n", err)
	}

	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Printf("tar FileInfoHeader err : %v\n", err)
	} else {
		fmt.Printf("hdr size is %v\n", hdr.Size)
		hdr.Size = int64(len(inserByte))
	}

	err = write.WriteHeader(hdr)
	if err != nil {
		fmt.Printf("write writeHeader err: %v\n", err)
	}

	ret ,err := write.Write(inserByte)
	if err != nil {
		fmt.Printf("write ./file.tar.gz err :%v\n", err)
	} else {
		fmt.Printf("write ./file.tar.gz success. return number is %d \n", ret)
	}

	err = write.Flush()
	if err != nil {
		fmt.Printf("write flush err : %v\n", err)
	}

	err = write.Close()
	if err != nil {
		fmt.Printf("write close err :%v\n", err)
	}

}

func PathExists(path string) (bool, error) {
	/*
	判断文件或文件夹是否存在
	如果返回的错误为nil,说明文件或文件夹存在
	如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	如果返回的错误为其它类型,则不确定是否在存在
	*/
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(dstName, srcName string) (writeen int64, err error) {
	src, err := os.Open(dstName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(srcName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
