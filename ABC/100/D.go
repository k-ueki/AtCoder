package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i], &y[i], &z[i])
	}

	nums := make([][]int, 8)
	for i := range nums {
		nums[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		count := 0
		for i := 0; i < 2; i++ {
			x[l] *= (-1)
			for j := 0; j < 2; j++ {
				y[l] *= (-1)
				for k := 0; k < 2; k++ {
					z[l] *= (-1)

					ans := x[l] + y[l] + z[l]
					nums[count][l] = ans
					count++
				}
			}
		}
	}

	max := 0
	for i := 0; i < 8; i++ {
		sort.Ints(nums[i])
		ans := 0
		for j := n - 1; j > n-1-m; j-- {
			ans += nums[i][j]
		}
		if ans > max {
			max = ans
		}
	}
	fmt.Println(max)
}
