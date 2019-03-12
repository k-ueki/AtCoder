package main

import (
	"fmt"
)

func resolve(str string) bool {

	var keyword = []string{"k", "e", "y", "e", "n", "c", "e"}
	typen := make(map[int]map[int]string)

	for i := 0; i < len(keyword); i++ {
		typen[i] = make(map[int]string)
		for j := 0; j <= i; j++ {
			typen[i][0] += keyword[j]
		}
		for k := i + 1; k <= len(keyword)-1; k++ {
			typen[i][1] += keyword[k]
		}
		if i == 6 {
			typen[i][1] = " "
		}
	}

	ans := false
	if str[:7] == "keyence" || str[len(str)-7:] == "keyence" {
		return true
	}
	for i := 0; i < 7; i++ {
		if str[:i+1] == typen[i][0] && str[len(str)-1-(7-i)+2:] == typen[i][1] {
			ans = true
		}

	}
	return ans
}
func main() {
	var str string
	fmt.Scanf("%s", &str)
	if resolve(str) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
