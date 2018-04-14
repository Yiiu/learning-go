package main

import "fmt"

type data struct {
	arr []int
}

func main() {
	arr := data{[]int{4, 2, 3, 1, 5, 6}}
	arr.bubbleSort()
	fmt.Println(arr)
}
// 冒泡
func (data *data) bubbleSort() {
	len := len(data.arr)
	for i := 0; i < len; i++ {
		for j := 1; j < len - i; j++ {
			if data.arr[j - 1] > data.arr[j] {
				new := data.arr[j]
				data.arr[j] = data.arr[j - 1]
				data.arr[j - 1] = new
			}
		}
	}
}