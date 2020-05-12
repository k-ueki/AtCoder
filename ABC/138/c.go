package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	v := make([]int, n)
	for i, _ := range v {
		fmt.Scan(&v[i])
	}

	sort.Ints(v)

	ans := float64(v[0])
	for i := 1; i < n; i++ {
		ans = calc(ans, float64(v[i]))
	}
	fmt.Println(ans)
}

func calc(x, y float64) float64 {
	return (x + y) / 2.0
}
