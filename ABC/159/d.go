package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			if i != j {
				tmp = append(tmp, a[j])
			}
		}
		fmt.Println(tmp)
	}
}
