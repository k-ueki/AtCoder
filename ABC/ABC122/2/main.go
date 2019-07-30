package main

import (
	"fmt"
)

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &q, &s)

	var l = make([]int, q)
	var r = make([]int, q)
	var str = make([]string, q)
	var ans = make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&l[i], &r[i])
		str[i] = s[l[i]-1 : r[i]]
	}
	i := 0
	for i < q {
		count := 0
		for j := 0; j < len(str[i]); j++ {
			if string(str[i][j]) == "A" && string(str[i][j+1]) == "C" {
				j++
				count++
			}
			if j == len(str[i])-2 {
				break
			}

		}
		ans[i] = count
		i++
	}
	for i := range ans {
		fmt.Println(ans[i])
	}
}
