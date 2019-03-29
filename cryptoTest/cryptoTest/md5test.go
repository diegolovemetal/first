package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	str := "abc123"

	data := []byte(str)
	has := md5.Sum(data)

	md5str1 := fmt.Sprintf("%x", has)

	fmt.Println(md5str1)

	w := md5.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	md5str2 := hex.EncodeToString(bw)
	fmt.Println(md5str2)
}
