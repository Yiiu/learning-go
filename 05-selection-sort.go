package main

import "fmt"

func main() {
	arr := data{[]int{4, 2, 3, 1, 5, 6}}
	arr.selectionSort()
	fmt.Println(arr)
}
// 选择
func (data *data) selectionSort() {
	len := len(data.arr)
	for i := 0; i < len; i++ {
		min := i
		for j := i; j < len; j++ {
			if data.arr[min] > data.arr[j] {
				min = j
			}
		}
		if min != i {
			tmp := data.arr[i]
			data.arr[i] = data.arr[min]
			data.arr[min] = tmp
		}
	}
}