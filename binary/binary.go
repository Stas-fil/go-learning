package main

import "fmt"

func binary_search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	numbers := []int{1, 3, 4, 6, 8, 9, 11, 12, 13, 16, 18, 20, 25, 28, 56, 59}
	target := 59

	index := binary_search(numbers, target)
	if index != -1 {
		fmt.Printf("Число %d найлено на позиции %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найлено\n", target)
	}
}
