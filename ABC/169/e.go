package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)

	min := float64(a[0]+b[0]) / 2
	max := float64(a[len(a)-1]+b[len(b)-1]) / 2
	fmt.Println(min, max)
	count := 0
	for i := min; i <= max; i += 0.5 {
		count++
	}
	fmt.Println("count", count)
}
