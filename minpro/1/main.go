package main

import (
	"fmt"
)

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)

	var bis int = 1
	turn := 0
	for turn < k {
		if b-a <= 2 {
			fmt.Println(k + 1)
			return
		} else {
			//fmt.Println("AA", bis, turn)
			if bis >= a && k-turn >= 2 {
				temp := bis / a
				//fmt.Println("temp", temp)
				bis -= temp * a
				//fmt.Println("bis", bis)
				bis += temp * b
				//fmt.Println("BIS", bis)
				turn += 2
			} else {
				//fmt.Println("b")
				bis++
				turn++
			}
		}
	}
	fmt.Println(bis)
}
