package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i, _ := range h {
		fmt.Scan(&h[i])
	}

	ans := 0
	temp := 0
	for i := 0; i < n-1; i++ {
		if h[i] >= h[i+1] {
			temp++
			if i == n-2 {
				ans = max(ans, temp)
			}
		} else {
			ans = max(ans, temp)
			temp = 0
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
