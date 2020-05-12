package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i, _ := range h {
		fmt.Scan(&h[i])
	}
}
