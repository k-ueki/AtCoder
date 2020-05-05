package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ans := make(map[int]int, n)
	for i := 2; i <= n; i++ {
		var temp int
		fmt.Scan(&temp)
		ans[temp]++
	}

	for i := 1; i <= n; i++ {
		fmt.Println(ans[i])
	}
}
