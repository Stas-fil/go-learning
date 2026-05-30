package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read all integers from stdin (supports spaces and newlines)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var arr []int
	for scanner.Scan() {
		token := scanner.Text()
		v, err := strconv.Atoi(token)
		if err != nil {
			continue
		}
		arr = append(arr, v)
	}

	if len(arr) <= 1 {
		fmt.Println("YES")
		return
	}

	isStrictlyIncreasing := true
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			isStrictlyIncreasing = false
			break
		}
	}

	if isStrictlyIncreasing {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
