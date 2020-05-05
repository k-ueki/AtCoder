package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	taxa := 8
	taxb := 10

	for i := 1; i < 10000; i++ {
		vala := i * taxa / 100
		valb := i * taxb / 100
		if vala == a && valb == b {
			fmt.Println(i)
			return
		}
		if i == 10000-1 {
			fmt.Println(-1)
		}
	}
}
