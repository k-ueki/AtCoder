package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		temp := 0
		fmt.Scan(&temp)
		a[i] = temp
	}

	temp := map[int]int{}
	val := []int{}
	for i, v := range a {
		temp[v] = i
		val = append(val, v)
	}
	sort.Ints(val)
	for _, v := range val {
		fmt.Printf("%d ", temp[v])
	}
}
