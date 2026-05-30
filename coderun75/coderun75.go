package main

import (
	"fmt"
)

func main() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	// Проверяем корректность входных данных
	if a < 1 || b < 1 || n < 0 || m < 0 || a > 1000 || b > 1000 || n > 1000 || m > 1000 {
		fmt.Println(-1)
		return
	}

	// Вычисляем время для каждого пути
	// Для n поездов: n минут стояния + (n-1) интервалов по a минут
	time1Min := n
	time1Max := n + (n-1)*a

	time2Min := m
	time2Max := m + (m-1)*b

	// Если один из путей не видел поездов, корректируем
	if n == 0 {
		time1Min, time1Max = 0, 0
	}
	if m == 0 {
		time2Min, time2Max = 0, 0
	}

	// Находим пересечение интервалов
	totalMin := max(time1Min, time2Min)
	totalMax := min(time1Max, time2Max)

	// Если пересечения нет, то решение невозможно
	if totalMin > totalMax {
		fmt.Println(-1)
	} else {
		fmt.Println(totalMin, totalMax)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
