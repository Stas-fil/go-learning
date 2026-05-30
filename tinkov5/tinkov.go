package main

import (
	"fmt"
	"sort"
)

func canCover(n int, intervals [][2]int) (string, int) {
	
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count, pos, i := 0, 1, 0

	for pos <= n {
		found := false
		maxR := 0

		
		for i < len(intervals) && intervals[i][0] <= pos {
			if intervals[i][1] > maxR {
				maxR = intervals[i][1]
			}
			i++
			found = true
		}

		
		if !found {
			return "No", 0
		}

		count++
		pos = maxR + 1
	}

	return "Yes", count
}

func main() {
	n := 5
	intervals := [][2]int{{1, 2}, {2, 3}, {3, 5}}

	res, num := canCover(n, intervals)
	fmt.Println(res)
	if res == "Yes" {
		fmt.Println(num)
	}
}
