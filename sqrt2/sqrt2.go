package main

import "fmt"

func sqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x/2
	var result int

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	result := sqrt(60)
	fmt.Println(result)
}
