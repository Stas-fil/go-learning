package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для представления отрезка
type Segment struct {
	l, r int
}

// Функция для проверки, можно ли восстановить сумму всех элементов
func canRestoreSum(n int, segments []Segment) (bool, int) {
	// Создаем граф смежности для префиксных сумм
	adj := make([][]int, n+1)

	// Добавляем ребра для каждого отрезка
	// Отрезок [l, r] означает, что мы знаем сумму S[r] - S[l-1]
	for _, seg := range segments {
		// Соединяем l-1 и r
		adj[seg.l-1] = append(adj[seg.l-1], seg.r)
		adj[seg.r] = append(adj[seg.r], seg.l-1)
	}

	// Проверяем связность от 0 до n
	visited := make([]bool, n+1)
	dfs(0, visited, adj)

	if !visited[n] {
		return false, 0
	}

	// Если граф связен, то можно восстановить сумму
	// Находим минимальное количество фактов для связности
	return true, findMinEdges(n, segments)
}

// Функция для проверки связности графа с помощью DFS
func dfs(node int, visited []bool, adj [][]int) {
	visited[node] = true
	for _, next := range adj[node] {
		if !visited[next] {
			dfs(next, visited, adj)
		}
	}
}

// Функция для нахождения минимального количества ребер для связности
func findMinEdges(n int, segments []Segment) int {
	// Используем алгоритм Крускала для нахождения MST
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var union func(int, int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}

	count := 0
	for _, seg := range segments {
		if find(seg.l-1) != find(seg.r) {
			union(seg.l-1, seg.r)
			count++
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем N и Q
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(line[0])
	q, _ := strconv.Atoi(line[1])

	// Читаем отрезки
	var segments []Segment
	for i := 0; i < q; i++ {
		scanner.Scan()
		fact := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(fact[0])
		r, _ := strconv.Atoi(fact[1])
		segments = append(segments, Segment{l, r})
	}

	// Проверяем, можно ли восстановить сумму
	possible, minFacts := canRestoreSum(n, segments)

	if possible {
		fmt.Println("Yes")
		fmt.Println(minFacts)
	} else {
		fmt.Println("No")
	}
}
