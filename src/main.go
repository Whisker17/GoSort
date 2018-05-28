package main

import "fmt"

func main() {
	datas := []int{2, 1, 5, 3, 4}
	datas = BubbleSort(datas)
	fmt.Println(datas)
}
