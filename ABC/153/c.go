package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i, _ := range h {
		fmt.Scan(&h[i])
	}

	sort.Ints(h)
	var temp []int

	for i := n - 1 - k; i >= 0; i-- {
		temp = append(temp, h[i])
	}
	ans := 0
	for _, v := range temp {
		ans += v
	}
	fmt.Println(ans)
}
