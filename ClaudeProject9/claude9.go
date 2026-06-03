package main

import "fmt"

func main() {

	add := func(a, b int) int { return a + b }
	sub := func(a, b int) int { return a - b }
	mul := func(a, b int) int { return a * b }
	fmt.Println(apply(3, 4, sub))
	fmt.Println(apply(3, 4, mul))
	fmt.Println(apply(3, 4, add))

}
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}
