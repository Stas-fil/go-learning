package main

import "fmt"

func calculator(operation string, a float64, b float64) float64 {
	switch operation {
	case "Сумма":
		return a + b
	case "Вычитание":
		return a - b
	case "Умножение":
		return a * b
	case "Деление":
		if b != 0 {
			return a / b
		}
		fmt.Println("Eroor")
		return 0
	default:
		fmt.Println("Eroor")
		return 0
	}

}

func main() {
	var operation string
	var a float64
	var b float64

	fmt.Println("Введите что хотите сделать:")
	fmt.Scan(&operation)

	fmt.Println("Введите первое число:")
	fmt.Scan(&a)

	fmt.Println("Введите второе число:")
	fmt.Scan(&b)

	var result float64 = calculator(operation, a, b)
	fmt.Printf(" Ваш ответ = %v\n", result)
}
