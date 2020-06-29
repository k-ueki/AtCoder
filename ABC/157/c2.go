package main

import (
	"fmt"
	"math"
	"strconv"
)

type (
	Inn struct {
		index int
		value int
	}
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	in := []Inn{}
	for i := 0; i < m; i++ {
		ind := 0
		val := 0
		fmt.Scan(&ind, &val)
		in = append(in, Inn{
			index: ind,
			value: val,
		})
	}

	for i := int(math.Pow(10, float64(n-1))); i < 1000; i++ {
		str := strconv.Itoa(i)
		for j, v := range in {
			if string(str[v.index-1]) != strconv.Itoa(v.value) {
				break
			}
			if j == len(in)-1 {
				fmt.Println(i)
				return
			}
		}
	}
	fmt.Println(-1)
}
