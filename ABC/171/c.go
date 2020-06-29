package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	str := []string{
		"z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y",
	}

	ans := ""
	var hoge []string
	for n != 0 {
		b := n % 26
		hoge = append(hoge, str[b])
		if b == 0 {
			n = n/26 - 1
			continue
		}
		n /= 26
	}
	for i := len(hoge) - 1; i >= 0; i-- {
		ans = fmt.Sprintf("%s%s", ans, hoge[i])
	}
	fmt.Println(ans)
}
