package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	temp := primeFactor(n)

	if n == 1 {
		fmt.Println(0)
		return
	}
	for i, v := range temp {
		for j := 1; j <= v; j++ {
			h := int(math.Pow(float64(i), float64(j)))
			if can(h) && n%h == 0 {
				n /= h
				count++
			}
			if n == 1 {
				fmt.Println(count)
				return
			}
		}
	}
	fmt.Println(count)

}

func can(n int) bool {
	if len(primeFactor(n)) == 1 {
		return true
	}
	return false
}

func primeFactor(n int) map[int]int {
	var res = make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res[i]++
			n /= i
		}
	}
	if n != 1 {
		res[n] = 1
	}
	return res
}
