package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	h := make([]int, n)
	temp := make([]int, n)

	for i, _ := range h {
		fmt.Scan(&h[i])
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		temp[a-1] = max(temp[a-1], h[b-1])
		temp[b-1] = max(temp[b-1], h[a-1])
	}

	ans := 0
	for i, v := range temp {
		if h[i] > v {
			ans++
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
