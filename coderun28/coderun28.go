package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	

}
