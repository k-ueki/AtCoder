package main

import "fmt"

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)
	c := make([]int, n)
	a := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
		a[i] = make([]int, m)
		for j, _ := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}

	for i, v := range a[0] {

	}

}
