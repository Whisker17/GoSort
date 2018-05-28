package main

import "fmt"

/*
选择排序是一种简单直观的排序算法，无论什么数据进去都是 O(n²) 的时间复杂度。所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧。

过程:
1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3.重复第二步，直到所有元素均排序完毕。
*/

func SelectionSort(datas []int) []int {
	len := len(datas)

	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if datas[min] > datas[j] {
				min = j
			}
		}
		datas[min], datas[i] = datas[i], datas[min]
	}
	return datas
}

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = SelectionSort(datas)
	fmt.Println(datas)
}
