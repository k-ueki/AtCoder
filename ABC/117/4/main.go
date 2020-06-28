package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func NextSlice(n int) []int {
	var mat []int
	in.Scan()
	temp := strings.Split(in.Text(), " ")
	for i := 0; i < n; i++ {
		li, _ := strconv.Atoi(temp[i])
		mat = append(mat, li)
	}
	return mat
}
func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	rep := NextSlice(k)

	fmt.Println(rep)
}
