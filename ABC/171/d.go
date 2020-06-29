package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	a := make([]int, n)
	count := make(map[int]int, n)
	ans := 0
	for i, _ := range a {
		fmt.Scan(&a[i])
		count[a[i]]++
		ans += a[i]
	}
	fmt.Scan(&q)
	b := make([]int, q)
	c := make([]int, q)
	for i, _ := range b {
		fmt.Scan(&b[i], &c[i])
	}

	for i := 0; i < q; i++ {
		ans -= b[i] * count[b[i]]
		ans += c[i] * (count[b[i]])
		tmp := count[b[i]]
		count[b[i]] = 0
		count[c[i]] += tmp
		fmt.Println(ans)
	}
}
