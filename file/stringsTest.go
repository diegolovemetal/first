package main

import (
	"fmt"
	"strconv"
)

func main() {
	////字符串s中是否包含substr，返回bool值
	//fmt.Println(strings.Contains("abc", "a"))
	////字符串链接，把slice a通过sep链接起来
	//s := []string{"food", "drink", "fruit"}
	//fmt.Println(strings.Join(s, ", "))
	////在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	//fmt.Println(strings.Index("chicken", "ken"))
	////重复s字符串count次，最后返回重复的字符串
	//fmt.Println("ba" + strings.Repeat("na", 2))
	////在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	//fmt.Println(strings.Replace("oi oi oi", "i", "h", -1))
	////把s字符串按照sep分割，返回slice
	//fmt.Printf("%q\n", strings.Split("a man is a kind of alive living", "a "))
	////在s字符串的头部和尾部去除cutset指定的字符串
	//fmt.Printf("[%q]", strings.Trim("aaa SASA aaa", "aaa"))
	////去除s字符串的空格符，并且按照空格分割返回slice
	//fmt.Printf("\nfields are: %q", strings.Fields("abc def ghi"))


	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	//str := make([]byte, 0, 100)
	//str = strconv.AppendInt(str, 15, 16)
	//str = strconv.AppendBool(str, false)
	//str = strconv.AppendQuote(str, "quote")
	//str = strconv.AppendQuoteRune(str, '汉')
	//fmt.Println(string(str))

	//e := strconv.Itoa(1023)
	//b := strconv.FormatFloat(123.23, 'g', 12, 64)
	//fmt.Println(b, e)

	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}