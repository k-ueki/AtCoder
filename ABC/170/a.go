package main

import "fmt"

func main() {
	x := make([]int, 5)
	for i, _ := range x {
		fmt.Scan(&x[i])
	}

	for i, v := range x {
		if v == 0 {
			fmt.Println(i + 1)
		}
	}
}
