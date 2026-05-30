package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	// Читаем стоимости обедов
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&costs[i])
	}
	
	// Вычисляем минимальную стоимость
	minCost := calculateMinCost(costs)
	fmt.Println(minCost)
}

func calculateMinCost(costs []int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	
	// dp[i] - минимальная стоимость для первых i дней
	dp := make([]int, n+1)
	dp[0] = 0
	
	for i := 1; i <= n; i++ {
		// Базовый случай: платим за текущий обед
		dp[i] = dp[i-1] + costs[i-1]
		
		// Проверяем, можем ли использовать купон
		// Купон получаем при покупке больше чем на 100 рублей
		if costs[i-1] > 100 {
			// Ищем лучший день для использования купона
			for j := i - 1; j >= 0; j-- {
				// Если можем использовать купон на день j
				if j > 0 && costs[j-1] > 100 {
					// Стоимость = стоимость до дня j-1 + стоимость дня j (без купона) + стоимость после дня i
					cost := dp[j-1] + costs[j-1] + sum(costs[j:i-1])
					if cost < dp[i] {
						dp[i] = cost
					}
				}
			}
		}
	}
	
	return dp[n]
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}