package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i])
	}

}
