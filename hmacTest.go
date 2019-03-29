package main

import (
	"crypto/hmac"
	"crypto/md5"
	sha12 "crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Md5(data string) string {
	//test := md5.New()
	//test.Write([]byte(data))
	//md5Data := md5.Sum([]byte(""))
	//return hex.EncodeToString(md5Data)
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}
func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func Sha1(data string) string {
	sha1 := sha12.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))

}

//func Hmac(key, data string) string {
//	hmac := hmac2.New(md5.New(), []byte(key))
//}
func main() {
	fmt.Println(Md5("hello"))
	fmt.Println(Hmac("key2", "hello"))
	fmt.Println(Sha1("hello"))
}
