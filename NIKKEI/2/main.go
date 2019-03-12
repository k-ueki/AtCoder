package main

import "fmt"

func main() {
	var n int
	var a, b, c string
	fmt.Scanf("%d\n%s\n%s\n%s\n", &n, &a, &b, &c)

	count := 0
	var temp string
	var temp2 string
	for i := 0; i < n; i++ {
		temp_count := 0
		temp = string(a[i])
		temp2 = string(b[i])
		if temp == string(b[i]) {
			temp_count++
		}
		if temp == string(c[i]) {
			temp_count++
		}

		if temp_count == 0 && temp2 == string(c[i]) {
			count++
			continue
		}
		if temp_count == 0 {
			count += 2
		} else if temp_count == 1 {
			count++
		} else if temp_count == 0 {
			count++
		}
	}
	fmt.Println(count)
}
