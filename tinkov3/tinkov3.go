package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)
	S, _ := reader.ReadString('\n')
	S = strings.TrimSpace(S)

	// Используем слайс для хранения последовательности
	sequence := []int{0}

	// Для быстрого поиска позиции числа i-1 используем map
	positions := make(map[int]int)
	positions[0] = 0

	for i := 1; i <= N; i++ {
		prevNum := i - 1
		prevPos := positions[prevNum]

		if S[i-1] == 'L' {
			// Вставляем слева от i-1 (в позицию prevPos)
			sequence = append(sequence[:prevPos], append([]int{i}, sequence[prevPos:]...)...)
			positions[i] = prevPos
			// Обновляем позиции всех чисел справа от вставленного
			for j := prevPos + 1; j < len(sequence); j++ {
				positions[sequence[j]] = j
			}
		} else {
			// Вставляем справа от i-1 (в позицию prevPos + 1)
			sequence = append(sequence[:prevPos+1], append([]int{i}, sequence[prevPos+1:]...)...)
			positions[i] = prevPos + 1
			// Обновляем позиции всех чисел справа от вставленного
			for j := prevPos + 2; j < len(sequence); j++ {
				positions[sequence[j]] = j
			}
		}
	}

	// Выводим результат
	for i, num := range sequence {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}
