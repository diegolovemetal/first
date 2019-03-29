package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	str := "你国人牛批"
	fmt.Printf("strting : %v\n", str)

	input := []byte(str)
	fmt.Printf("[]byte : %v\n", input)

	encodeStr := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("[]byte :%v\n",encodeStr)

	decodeBytes, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("decode base64 :%v\n", string(decodeBytes))

	fmt.Println()
	//url需要URLEncoding
	urlencode := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Printf("urlencode :%v\n", urlencode)

	urldecode, err := base64.URLEncoding.DecodeString(urlencode)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("urldecode %v\n", string(urldecode))
}
