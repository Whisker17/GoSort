package main

import "fmt"

/*
归并排序（Merge sort）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：
1.自上而下的递归（所有递归的方法都可以用迭代重写，所以就有了第 2 种方法）；
2.自下而上的迭代；
和选择排序一样，归并排序的性能不受输入数据的影响，但表现比选择排序好的多，因为始终都是 O(nlogn) 的时间复杂度。代价是需要额外的内存空间。

步骤：
1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
4.重复步骤 3 直到某一指针达到序列尾；
5.将另一序列剩下的所有元素直接复制到合并序列尾。
*/

func MergeSort(datas []int) []int {
	len := len(datas)
	if len < 2 {
		return datas
	}
	mid := len / 2
	left := datas[0:mid]
	right := datas[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}

	return res
}

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = MergeSort(datas)
	fmt.Println(datas)
}
