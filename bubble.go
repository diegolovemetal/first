package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 6, 8}
	bubleSort(sli)
	fmt.Println(sli)

}

func bubleSort(sli []int) []int {
	len := len(sli)
	for i := 1; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	return sli
}