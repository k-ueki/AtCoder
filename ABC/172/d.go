package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := make(map[int]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			count[j]++
		}
	}

	ans := big.NewInt(0)
	for i := 1; i <= n; i++ {
		tmp := big.NewInt(int64(i) * int64(count[i]))
		ans.Add(ans, tmp)
	}
	fmt.Println(ans)
}
