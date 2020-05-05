package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i, _ := range x {
		fmt.Scan(&x[i])
	}

	ans := 1000000000000
	for i := 1; i <= 100; i++ {
		temp := 0
		for _, v := range x {
			temp += (i - v) * (i - v)
		}
		ans = min(temp, ans)
	}
	fmt.Println(ans)

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
