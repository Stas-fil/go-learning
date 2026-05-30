package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := make([]int, 3)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	fmt.Println(arr[1])
}
