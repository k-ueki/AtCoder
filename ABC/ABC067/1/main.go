package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func NextSlices(n int) []int {
	rep := make([]int, n)
	in.Scan()
	sp := strings.Split(in.Text(), " ")
	temp, _ := strconv.Atoi(sp)
	rep[i] = temp
	return rep

}
func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	mat := NextSlices(n)
	fmt.Println(mat)
}
