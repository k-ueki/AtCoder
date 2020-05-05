package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i, _ := range s {
		fmt.Scan(&s[i])
	}

	ma := make(map[string]int)
	for _, v := range s {
		ma[v]++
	}

	max := 0
	for _, v := range ma {
		max = Max(v, max)
	}

	var ans []string
	for i, v := range ma {
		if v == max {
			ans = append(ans, i)
		}
	}
	sort.Strings(ans)
	for _, v := range ans {
		fmt.Println(v)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
