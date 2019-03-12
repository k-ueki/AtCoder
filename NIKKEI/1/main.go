package main

import "fmt"

func main() {

	var n, a, b int
	fmt.Scanf("%d %d %d\n", &n, &a, &b)

	var max, min int
	if a+b <= n {
		if a > b {
			max = b
			min = 0
		} else {
			max = a
			min = 0
		}
	} else {
		if a > b {
			max = b
			min = a - n + b
		} else {
			max = a
			min = a - n + b
		}
	}
	fmt.Println(max, min)
}
