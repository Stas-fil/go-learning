package main

import (
	"fmt"
	"math"
)

func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}
func main() {
	var radius float64 = 5
	area := calculateArea(radius)
	fmt.Printf("The area of circle with radius %.2f = %.2f\n", radius, area)
}
