package main

import "fmt"

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&a[i], &b[i], &c[i], &d[i])
	}

}
