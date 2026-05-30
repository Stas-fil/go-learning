package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 7
	y := 2.5
	fmt.Println(float64(x) + y)
	fmt.Println(x + int(y))
	n := 65
	fmt.Println(string(n))
	fmt.Println(strconv.Itoa(n))

}
