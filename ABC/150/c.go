package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	q := make([]int, n)
	for i, _ := range p {
		fmt.Scan(&p[i])
	}
	for i, _ := range q {
		fmt.Scan(&q[i])
	}

}
