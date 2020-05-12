package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	ans := 0
	for i := 1; i <= k; i++ {
		if i >= 1 && i <= a {
			ans++
		}
		if a+b < i && i <= k {
			ans--
		}
	}
	fmt.Println(ans)
}
