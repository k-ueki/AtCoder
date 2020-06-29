package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	in := make(map[int]string, m)
	for i := 0; i < m; i++ {
		ind := 0
		val := ""
		fmt.Scan(&ind, &val)
		if ind == 1 && val == "0" {
			fmt.Println(-1)
		}
		if in[ind] != "0" && in[ind] != val {
			fmt.Println(-1)
			return
		}
		in[ind] = val
	}

	for i := 122; i < int(math.Pow(10, float64(n))); i++ {
		str := strconv.Itoa(i)
		fmt.Println(i)
		fmt.Println(string(str[0]))
		for i, v := range str {
			if string(v) != in[i+1] {
				break
			}
			if i == len(str)-1 {
				fmt.Println(i)
				return
			}
		}
	}
}

func keta(i int) int {
	hoge := 0
	for i != 0 {
		i /= 10
		hoge++
	}
	return hoge
}
