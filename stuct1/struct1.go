package main

import "fmt"

func main() {
	var arr [3]int = [3]int{5, 7, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr2 := []int{}
	arr2 = append(arr2, 4, 5, 6)
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	arr3 := map[string]int{"apple": 1, "banana": 2, "orange": 3, "melon": 4}
	fmt.Println(arr3["apple"])
	for k, v := range arr3 {
		if v == 3 {
			fmt.Println("Key:", k)
			break
		}
	}

	
}
