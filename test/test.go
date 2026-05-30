package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var number int
	count := 0
	numberComp := rand.Intn(100) + 1

	for {
		fmt.Print("Назовите ваше число: ")
		fmt.Scanln(&number)

		count++

		if number == numberComp {
			fmt.Printf("🎉 Вы угадали! Компьютер загадал: %d\n", numberComp)
			break
		} else if number < numberComp {
			fmt.Printf("🙁 Ваше число меньше загаданного.\n")
		} else {
			fmt.Printf("🙁 Ваше число больше загаданного.\n")
		}
	}

	fmt.Printf("Вы сделали %d попыток.\n", count)
}
