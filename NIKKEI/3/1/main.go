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

//1:一致
func judge_same(a, b []int, n int) int {
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return 0
		}
	}
	return 1
}
func judge_each(a, b []int, n int) (ansa, ansb int) {
	tempa := a[0]
	tempb := b[0]
	ansa = 1
	ansb = 1
	for i := 1; i < n; i++ {
		if a[i] != tempa {
			ansa = 0
		}
		if b[i] != tempb {
			ansb = 0
		}
		if ansa == 0 && ansb == 0 {
			break
		}
	}
	return ansa, ansb
}
func calc(a, b []int, n int) (sumT, sumA int) {
	max_c := 0
	sumT, sumA = 0, 0
	//一致の場合
	for i := 0; i < n; i++ {
		fmt.Println(i)
		max := 0
		for j := 0; j < len(a); j++ {
			if max < a[j] {
				max = a[j]
				max_c = j
			}
		}
		fmt.Println("MAX", max, "MAX_C", max_c)

		if i%2 == 0 {
			fmt.Println("T")
			sumT += max
		} else {
			fmt.Println("A")

			sumA += max
		}
		a = append(a[:max_c], a[max_c+1:]...)
		b = append(b[:max_c], b[max_c+1:]...)
		fmt.Println(a, b)
		fmt.Println(sumT, sumA)
	}
	return sumT, sumA
}
func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a, b := nextMaps(n)

	var flag bool
	sT, sA := 0, 0
	fa, fb := judge_each(a, b, n)
	fmt.Println("fa,fb", fa, fb)

	if judge_same(a, b, n) == 1 {
		flag = true
	} else {
		flag = false
	}
	if flag {
		sT, sA = calc(a, b, n)

	}
	fmt.Println(sT - sA)
}
