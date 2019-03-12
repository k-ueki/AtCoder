package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func nextMaps(num int) (mat_a, mat_b []int) {
	in.Split(bufio.ScanWords)

	mat_a = make([]int, num)
	mat_b = make([]int, num)

	for i := 0; i < num; i++ {
		//fmt.Fscanf(in, "%d ", &mat1[i])
		in.Scan()
		fmt.Sscanf(in.Text(), "%d", &mat_a[i])
	}
	for i := 0; i < num; i++ {
		//fmt.Fscanf(in, "%d ", &mat2[i])
		in.Scan()
		fmt.Sscanf(in.Text(), "%d", &mat_b[i])
	}
	return mat_a, mat_b
}
func sumMat(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
func resolve(as, bs []int, N int) (err error, ans int) {
	sumA := sumMat(as)
	sumB := sumMat(bs)

	if sumA < sumB {
		err = errors.New("fail")
		return err, -1
	} else {
		tempSumA := 0
		tempSumB := 0
		count := 0
		for i := 0; i < N; i++ {
			if as[i] < bs[i] {
				tempSumA += as[i]
				tempSumB += bs[i]
				count++
			}
			if tempSumA >= tempSumB {
			}
		}
		return nil, 99
	}
}
func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	a, b := nextMaps(N)
	e, ans := resolve(a, b, N)

	if e == nil {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}
