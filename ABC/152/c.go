package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i, _ := range p {
		fmt.Scan(&p[i])
	}

	ans := 0
	min := 9999999999999
	for i := 0; i < n; i++ {
		min = Min(min, p[i])
		if p[i] <= min {
			ans++
		}
	}
	fmt.Println(ans)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
