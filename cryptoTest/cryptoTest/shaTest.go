package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	str := "掉泪老牟"

	data := []byte(str)
	has := sha1.Sum(data)
	sha1str1 := fmt.Sprintf("%x", has)

	fmt.Println(sha1str1)

	w := sha1.New()
	io.WriteString(w, str)
	bw := w.Sum(nil) //w.sum(nil)j将W的hash转成[]byte格式

	shastr2	:= hex.EncodeToString(bw)
	fmt.Println(shastr2)

	w = sha256.New()
	io.WriteString(w, str)
	bw = w.Sum(nil)

	sha256str1 := hex.EncodeToString(bw)
	fmt.Println(sha256str1)

	w = sha512.New()
	io.WriteString(w, str)
	bw = w.Sum(nil)

	sha512str := hex.EncodeToString(bw)
	fmt.Println(sha512str)

}
