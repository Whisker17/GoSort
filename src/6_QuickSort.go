package main

import "fmt"

/*
快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
快速排序又是一种分而治之思想在排序算法上的典型应用。本质上来看，快速排序应该算是在冒泡排序基础上的递归分治法。

步骤：
1.从数列中挑出一个元素，称为 “基准”（pivot）;
2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
  在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/

func QuickSort(datas []int) []int {
	return _quickSort(datas, 0, len(datas)-1)
}

func _quickSort(datas []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(datas, left, right)
		_quickSort(datas, left, partitionIndex-1)
		_quickSort(datas, partitionIndex+1, right)
	}
	return datas
}

func partition(datas []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if datas[i] < datas[pivot] {
			datas[i], datas[index] = datas[index], datas[i]
			index++
		}
	}
	datas[pivot], datas[index-1] = datas[index-1], datas[pivot]
	return index - 1
}

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = QuickSort(datas)
	fmt.Println(datas)
}
