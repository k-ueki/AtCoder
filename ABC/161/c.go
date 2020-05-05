package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	n %= k
	min := Min(n, abs(n-k))
	fmt.Println(min)
}

func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
