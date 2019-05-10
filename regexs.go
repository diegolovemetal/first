package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := " I am learning GO language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	//符合正则规则第一个
	one := re.Find([]byte(a))
	fmt.Println("FIND:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FINDALL", all)

	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FINDINDEX", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FINDALLINDEX", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FINDSUBMATCH", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
