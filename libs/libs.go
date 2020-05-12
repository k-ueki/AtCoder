package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

//"""exits?"""
//=============入力===============

//引数nの長さのスライスを生成する
//i1 i2 i3 ... in
func nextSlice(n int) []int {
	var mat []int
	in.Scan()
	temp := strings.Split(in.Text(), " ")
	for i := 0; i < n; i++ {
		li, _ := strconv.Atoi(temp[i])
		mat = append(mat, li)
	}
	return mat
}

//こんな面倒なことしなくても
//var mat = make([]int,n)
//for i := range mat{
//	fmt.Scan(&mat[i])
//}

//引数nの長さのスライスを生成
//a1 b1
//a2 b2....
func nextSlices(n int) (a, b []int) {
	for i := 0; i < n; i++ {
		in.Scan()
		temp := strings.Split(in.Text(), " ")
		li1, _ := strconv.Atoi(temp[0])
		li2, _ := strconv.Atoi(temp[1])
		a = append(a, li1)
		b = append(b, li2)
	}
	//これだけでOK
	//for i:=0;i<n;i++{
	//fmt.Scan(&a[i],&b[i])
	//}
	return a, b
}

//====================================
//WORD分割(一文字ずつ)
func separateWords(str string) []string {
	temp := strings.Split(str, "")
	return temp
}

//最大値
//sliceから
func max(num []int) int {
	max := 0
	for i := 0; i < len(num); i++ {
		if max < num[i] {
			max = num[i]
		}
	}
	return max
}

//最小値
//from slice
func min(num []int) int {
	min := 9999999
	for i := 0; i < len(num); i++ {
		if min > num[i] {
			min = num[i]
		}
	}
	return min
}

//最小公倍数
func lcm(i, j int) int {
	return i * j / gcd(i, j)
}

//最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//==========階乗、順列、組合せ、等==============

//順列
func permutation(n int, k int) int {
	res := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			res *= (n - i)
		}
	} else if k > n {
		res = 0
	}
	return res
}

//階乗
func factorial(n int) int {
	return permutation(n, n-1)
}

//組合せ
func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

//重複組合せ
func combRepete(n int, k int) int {
	return combination(n+k-1, k)
}

//========================================================

//==========その他===============

//素数判定
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//素因数分解
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

func main() {
	//	var str string
	//	fmt.Scanf("%s\n", &str)
	//	word := SeparateWords(str)
	//	fmt.Println(word[3])
	//fmt.Println(primeFactor(18))
}
