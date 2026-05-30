package main

import (
	"fmt"
	"math/rand"
)

var distance = 1222000000

const HourPerDay = 24

func main() {
	Carrier := " "
	TripeType := " "
	fmt.Printf("Titan Express Booking\n")
	fmt.Printf("%-18v %-16v %-14v %-20v\n", "Carrier", "Days", "Trip Type", "Price(million$)")
	fmt.Printf("=======================================================================\n")
	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			Carrier = "Blue Horizont"
		case 1:
			Carrier = "Galactic Express"
		case 2:
			Carrier = "Neptune Lines"
		}
		speed := rand.Intn(50000) + 50001
		Days := distance / speed / HourPerDay
		price := (speed / 1000) + 30
		if rand.Intn(2) == 1 {
			TripeType = "Round-trip"
			price *= 2
		} else {
			TripeType = "One-way"
		}
		fmt.Printf("%-18v %-16v %-14v %-20v\n", Carrier, Days, TripeType, price)
	}
}
