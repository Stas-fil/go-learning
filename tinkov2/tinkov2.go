package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const ALPHABET = 26
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)
	n := len(s)

	cntL := make([]int, ALPHABET)

	cntR := make([]int, ALPHABET)

	for i := 0; i < n; i++ {
		cntR[s[i]-'a']++
	}

	var answer int64

	for j := 0; j < n; j++ {
		current := s[j] - 'a'
		cntR[current]--

		sumL := 0
		for i := 0; i <= int(current); i++ {
			sumL += cntL[i]
		}

		sumR := 0
		for i := int(current); i < ALPHABET; i++ {
			sumR += cntR[i]
		}

		answer += int64(sumL) * int64(sumR)
		cntL[current]++
	}

	fmt.Println(answer)
}
