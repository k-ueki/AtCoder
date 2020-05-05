package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	set := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		set[temp] = struct{}{}
	}
	fmt.Println(len(set))
}
