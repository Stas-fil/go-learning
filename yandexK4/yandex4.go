package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Уравнение √(ax + b) = c эквивалентно ax + b = c²
	// при условии, что c ≥ 0 и ax + b ≥ 0

	// Если c < 0, то решений нет, так как √(ax + b) ≥ 0
	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	// Получаем уравнение ax + b = c²
	// ax = c² - b
	// x = (c² - b) / a

	rightSide := c*c - b

	// Если a = 0
	if a == 0 {
		// Уравнение становится 0*x + b = c²
		// b = c²
		if b == c*c {
			// Любое x является решением
			fmt.Println("MANY SOLUTIONS")
		} else {
			// Решений нет
			fmt.Println("NO SOLUTION")
		}
		return
	}

	// Если a ≠ 0, то x = (c² - b) / a
	x := rightSide / a

	// Проверяем, что x - целое число
	if a*x != rightSide {
		fmt.Println("NO SOLUTION")
		return
	}

	// Проверяем, что подкоренное выражение неотрицательно
	if a*x+b < 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	fmt.Println(x)
}
