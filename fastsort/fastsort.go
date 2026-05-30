package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]

	var left, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
func main() {
	data := []int{10, 6, 3, 8, 9, 4, 12, 11, 50}
	sorted := quickSort(data)
	fmt.Println(sorted)
}
