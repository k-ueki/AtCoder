package main

import "fmt"

func main() {
	var a, n, k, q int
	fmt.Scan(&n, &k, &q)
	temp := make(map[int]int, n)
	for i := 0; i < q; i++ {
		fmt.Scan(&a)
		temp[a]++
	}

	for i := 1; i <= n; i++ {
		if q-temp[i] < k {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
