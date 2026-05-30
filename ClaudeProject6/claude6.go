package main

import "fmt"

func main() {
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b

	}
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}
}
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
