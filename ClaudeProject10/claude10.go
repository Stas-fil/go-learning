package main

import "fmt"

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5
	defer fmt.Println(a)
	defer fmt.Println(b)
	defer fmt.Println(c)
	defer fmt.Println(d)
	defer fmt.Println(e)

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
