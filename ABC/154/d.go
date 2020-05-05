package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	p := make([]int, n)
	for i, _ := range p {
		fmt.Scan(&p[i])
	}

	max := 0
	for i := 0; i < n; i++ {

	}
}

func expect(n int) float64 {
	var resp float64
	for i := 1; i <= n; i++ {
		resp += float64(i) / float64(n)
	}
	return resp
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
