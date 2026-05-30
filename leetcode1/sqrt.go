package main

import "fmt"

func sqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 0, x
	t := x
	var mid int
	for {
		mid = (left + right) / 2
		if mid == t || mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid
		} else {
			right = mid
		}
		t = mid

	}

}

func main() {
	result := sqrt(120)
	fmt.Println(result)
}
