package main

import (
	"fmt"
	"strconv"
)

func allDigitalSame(n int) bool {
	s := strconv.Itoa(n)
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	count := 0
	for i := l; i <= r; i++ {
		if allDigitalSame(i) {
			count++
		}
	}
	fmt.Println(count)
}
