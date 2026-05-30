package main

import "fmt"

func main() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Friday, Sunday)

	const Pi = 3.14159
	r := 5.0
	fmt.Println(2 * Pi * r)
}
