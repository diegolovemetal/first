package main

import "fmt"

var sli = []int {1,232,2421,122,3213,533,23,2132}

func insertSort(sli []int) []int {
	len := len(sli)
	for i:=0; i<len; i++ {
		tmp := sli[i]
		//内层循环控制
		for j:=i-1; j>=0; j-- {
			//发现插入的元素要小一些，将后面的元素和前面元素互换
			if tmp < sli[j] {
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				break
			}
		}
	}
	return  sli
}

func main() {
	res := insertSort(sli)
	fmt.Println(res)
}
