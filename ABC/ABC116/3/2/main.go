package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func NewSlice(n int) []int {
	rep := make([]int, n)
	in.Scan()
	ele := strings.Split(in.Text(), " ")
	for i := 0; i < n; i++ {
		n, _ := strconv.Atoi(ele[i])
		rep[i] = n
	}
	return rep
}
func iszero(num []int) bool {
	for _, v := range num {
		if v == 0 {
			return true
		}
	}
	return false
}
func resolve(num []int) {
	count := 0
	if iszero(num) {
	} else {
		fmt.Println("HEHEHE")
		for i := 0; i < len(num); i++ {
			num[i] = num[i] - 1
		}
		fmt.Println(num, count)
	}
}
func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	rep := NewSlice(n)
	resolve(rep)
	fmt.Println(rep)
}
