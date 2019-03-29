package main

import "fmt"

var sli =[]int {1,23,12,32,65,2,63}

//双重循环控制，外层控制轮数，内层控制比较次数
func selectSort(sli []int) []int {
	len := len(sli)
	for i:=0; i<len-1; i++ {
		k := i	//假设最小值的位置
		for j:=i+1; j<len; j++ {
			if sli[k] > sli[j] { //sli[k]为当前已知最小值
				k = j	//比较并记录下更小值的位置；在下一次比较时采用已知最小值
			}
		}
		//已经确定为最小值的位置，保存到K中，如果发现当前假设的值与i不相同，位置互换
		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}
	}
	return sli
}

func main() {
	res := selectSort(sli)
	fmt.Println(res)
}
