package main

import "fmt"

/*
堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：
  即子结点的键值或索引总是小于（或者大于）它的父节点。堆排序可以说是一种利用堆的概念来排序的选择排序。分为两种方法：
1.大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
2.小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；
堆排序的平均时间复杂度为 Ο(nlogn)。

步骤：
1.创建一个堆 H[0……n-1]；
2.把堆首（最大值）和堆尾互换；
3.把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
4.重复步骤 2，直到堆的尺寸为 1。
*/

func HeapSort(datas []int) []int {
	len := len(datas)
	buildMaxHeap(datas, len)
	for i := len - 1; i >= 0; i-- {
		datas[0], datas[i] = datas[i], datas[0]
		len--
		heapify(datas, 0, len)
	}
	return datas
}

func buildMaxHeap(datas []int, len int) {
	for i := len / 2; i >= 0; i-- {
		heapify(datas, i, len)
	}
}

func heapify(datas []int, i, len int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < len && datas[left] > datas[largest] {
		largest = left
	}
	if right < len && datas[right] > datas[largest] {
		largest = right
	}
	if largest != i {
		datas[i], datas[largest] = datas[largest], datas[i]
		heapify(datas, largest, len)
	}
}

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = HeapSort(datas)
	fmt.Println(datas)
}
