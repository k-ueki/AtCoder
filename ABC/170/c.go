package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	if n == 0 {
		fmt.Println(x)
		return
	}
	p := make([]int, n)
	for i, _ := range p {
		fmt.Scan(&p[i])
	}

	sort.Ints(p)

	var tmp []int
	for i := -10; i < 1001; i++ {
		for j := 0; j < n; j++ {
			if i == p[j] {
				break
			}
			if j == n-1 {
				tmp = append(tmp, i)
			}
		}
	}

	ans := 100100100100
	val := 0
	for _, v := range tmp {
		dist := abs(v - x)
		if dist < ans {
			ans = dist
			val = v
		}
	}
	fmt.Println(val)
}

func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}
