package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	sta := make(map[int]int, a[len(a)-1]+1)
	for _, v := range a {
		if sta[v] != 0 {
			sta[v] = 2
			continue
		}
		for i := v; i <= a[len(a)-1]; i += v {
			sta[i]++
		}
	}

	ans := 0
	for _, v := range a {
		if sta[v] == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
