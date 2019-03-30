package main

import (
	"fmt"
	"github.com/axgle/mahonia"
)

func ConvertTostring(src, srcCode, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult),true)
	rusult := string(cdata)
	return rusult
}

func main() {
	str := "����ת������"
	str = ConvertTostring(str, "gbk", "utf-8")
	fmt.Println(str)
}