package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 0
	for sum < 2500 {
		switch rand.Intn(3) {
		case 0:
			sum += 5
		case 1:
			sum += 10
		case 2:
			sum += 25
		}
		dollars := sum / 100
		cents := sum % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)
	}
}
