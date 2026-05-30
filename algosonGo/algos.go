package main

import "fmt"

func recursive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + recursive(arr[1:])
}
func main() {
	nums := []int{3, 4, 1, 5, 6, 3, 5, 2, 4, 9}
	result := recursive(nums)
	fmt.Println(result)
}
