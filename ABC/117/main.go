package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, m int
	fmt.scan(&n, &m)

	var rep = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.scan(&rep[i])
	}
	sort.ints(rep)

	var deff []int
	for i := 0; i < len(rep)-1; i++ {
		temp := math.abs(float64(rep[i]) - float64(rep[i+1]))
		deff = append(deff, int(temp))
	}
	sort.ints(deff)

	ans := 0
	for i := 0; i <= len(deff)-n; i++ {
		ans += deff[i]
	}
	fmt.println(ans)
}
