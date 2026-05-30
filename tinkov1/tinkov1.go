package main

import (
	"fmt"
	"log"
)

func main() {
	var n, m int
	fmt.Print("Введите номер текущего дня месяца и число записи две недели назад (например: 15 1): ")
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		log.Fatal("Ошибка: Введите два целых числа.")
	}

	if n < 1 || n > 31 || m < 1 || m > 31 {
		log.Fatal("Ошибка: числа должны быть от 1 до 31.")
	}

	// Прошлое воскресенье было 7 дней назад
	previousSunday := n - 7

	// Выводим результат
	fmt.Printf("Запись за прошлое воскресенье нужно подписать числом: %d\n", previousSunday)
}
