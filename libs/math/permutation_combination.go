package math

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

func makePermutations(i int) []int {

}
