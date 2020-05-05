package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]int, m)
	c := make([]int, m)
	for i, _ := range s {
		fmt.Scan(&s[i], &c[i])
	}
}
