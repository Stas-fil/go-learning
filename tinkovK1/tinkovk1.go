package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)
	if D > B {
		value := ((D - B) * C) + A
		fmt.Println(value)
	} else {
		fmt.Println(A)
	}
}
