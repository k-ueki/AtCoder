package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func nextMaps(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		sp := strings.Split(in.Text(), " ")
		li, _ := strconv.Atoi(sp[0])
		li2, _ := strconv.Atoi(sp[1])
		a[i] = li
		b[i] = li2
	}
	return a, b
}
func judge_a_same(a []int) bool {
	temp := a[0]
	for i := 1; i < len(a); i++ {
		if temp != a[i] {
			return false
		}
	}
	return true
}
func resolve(a, b []int, n int) {
	ans := 0
	max_c := 0

	if judge_a_same(a) {
		temp := a[0]
		for i := 0; i < n; i++ {
			max := 0
			for j := 0; j < len(b); j++ {
				if max < b[j] {
					max = b[j]
					max_c = j
				}
			}
			if i%2 == 0 {
				ans += temp
			} else {
				ans -= b[max_c]
			}
			a = append(a[:max_c], a[max_c+1:]...)
			b = append(b[:max_c], b[max_c+1:]...)
		}
	} else {
		for i := 0; i < n; i++ {
			max := 0
			if i%2 == 0 {
				for j := 0; j < len(a); j++ {
					if max < a[j] {
						max = a[j]
						max_c = j
					}
				}
				ans += a[max_c]
			} else {
				for j := 0; j < len(b); j++ {
					if max < b[j] {
						max = b[j]
						max_c = j
					}
				}
				ans -= b[max_c]
			}
			a = append(a[:max_c], a[max_c+1:]...)
			b = append(b[:max_c], b[max_c+1:]...)
		}
	}
	fmt.Println(ans)
}
func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	a, b := nextMaps(n)

	resolve(a, b, n)
}
