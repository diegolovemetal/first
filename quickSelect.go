package main

import "fmt"

var sli = []int{1,213,212,321,51,123,332,46,231}

func quickSort(sli []int) []int {
	len := len(sli)
	//首先判断是否继续进行
	if len < 1{
		return sli
	}
	//第一个元素或者最后一个元素作为基准元素
	base_num := sli[0]
	//遍历除了基准元素以外所有元素，按照大小关系放入左右两个切片

	left_sli := []int {}
	right_sli := []int {}

	for i:=1; i<len; i++ {
		if base_num > sli[i]{
			left_sli =append(left_sli, sli[i])
		} else {
			right_sli = append(right_sli, sli[i])
		}
	}
	//分别对左右两个切片进行相同的排序，递归调用这个函数
	left_sli = quickSort(left_sli)
	right_sli = quickSort(right_sli)

	//合并
	left_sli = append(left_sli, base_num)
	return append(left_sli, right_sli...)

}

func main() {
	res := quickSort(sli)
	fmt.Println(res)
}
