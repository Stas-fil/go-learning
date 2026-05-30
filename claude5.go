package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}
}
