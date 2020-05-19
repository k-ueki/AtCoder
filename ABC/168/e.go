package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i], &b[i])
	}

	fmt.Println(n)
	temp := make(map[int][]int)
	for i := 0; i < (1 << uint(n)); i++ {
		count := 0
		tempc := 0
		fmt.Println(i)
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				tempc++
				if tempc > 2 {

				}
				fmt.Println("hoge", j)
			}
		}
		fmt.Println()
	}

}
