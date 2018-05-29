package main

import "fmt"

/*
计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
*/

func CountingSort(datas []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	len := len(datas)

	for i := 0; i < len; i++ {
		bucket[datas[i]]++
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			bucket[j]--
			datas[sortedIndex] = j
			sortedIndex++
		}
	}

	return datas
}

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = CountingSort(datas, 5)
	fmt.Println(datas)
}
