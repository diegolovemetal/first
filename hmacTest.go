package main

import (
	md52 "crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(data string) string {
	md5 := md52.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum(nil)
	return hex.EncodeToString(md5Data)
}

func main() {
	fmt.Println(Md5("hello"))
}