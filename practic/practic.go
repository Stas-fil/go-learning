package main

import (
	"fmt"
	"math/rand"
)

var Distance = 62100000

const secondPerDay = 86400

func main() {
	fmt.Printf("Spaceline         Days Trip type Price\n")
	fmt.Printf("======================================\n")
	Spaceline := " "
	Type := " "

	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			Spaceline = "Space Adventures"
		case 1:
			Spaceline = "SpaceX"
		case 2:
			Spaceline = "Virgin Galactic"
		}

		Speed := rand.Intn(15) + 16
		Days := (Distance / Speed) / secondPerDay
		Price := 20 + Days
		if rand.Intn(2) == 1 {
			Type = "Round-trip"
			Price *= 2
		} else {
			Type = "One-way"
		}
		fmt.Printf("%-16v %4v %-10v $%4v\n", Spaceline, Days, Type, Price)
	}
}
