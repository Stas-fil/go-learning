package main

import (
	"fmt"
	"math/rand"
)

func main() {
	my := 18

	for {
		x := rand.Intn(20)
		fmt.Printf("Число компьютера = %v\n", x)
		if x > my {
			fmt.Printf("Число компьютера больше вашего: %v > %v\n", x, my)
		} else if x < my {
			fmt.Printf("Число компьютера меньше вашего: %v < %v\n", x, my)
		} else {
			fmt.Printf("Компьютер угадал ваше число: %v = %v\n", x, my)
			break
		}
	}
}
