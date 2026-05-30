package main

import "fmt"

func main() {
	var str string = "Hello"
	number := 10
	var tr bool = true
	float := 0.5
	fmt.Printf("%T %v | %T %v | %T %v | %T %v\n",
		str, str,
		number, number,
		tr, tr,
		float, float,
	)

}
