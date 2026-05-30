package main

import (
	"fmt"
)

func main() {
	var n int
	var k int64
	fmt.Scan(&n, &k)

	arr := make([]int64, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	// O(n) алгоритм с префиксными суммами
	prefixCount := make(map[int64]int64)
	prefixCount[0] = 1 // пустая сумма в начале

	var sum int64
	var count int64

	for i := range arr {
		sum += arr[i]
		// Ищем количество предыдущих префиксов, которые дают sum - k
		count += prefixCount[sum-k]
		prefixCount[sum]++
	}

	fmt.Println(count)
}
