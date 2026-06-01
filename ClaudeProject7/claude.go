package main

import "fmt"

func main() {
	fmt.Printf("%3s", "")
	for s := 1; s < 10; s++ {
		fmt.Printf("%3d", s)

	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		fmt.Printf("%3d", i)
		for j := 1; j < 10; j++ {
			fmt.Printf("%3d", j*i)
		}
		fmt.Println()

	}
}
