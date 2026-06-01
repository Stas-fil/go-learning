package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{12, 34, 62, 13, 64, 12, 576, 34, 1, 553, 34, 76}
	min, max, err := minMax(nums)
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("min:", min)
		fmt.Println("max:", max)
	}

}
func minMax(nums []int) (int, int, error) {
	min, max := nums[0], nums[0]
	if len(nums) == 0 {
		return 0, 0, errors.New("Empty")
	}
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max, nil
}
