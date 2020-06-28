package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ans := 0
	for i, v := range s {
		if string(v) != string(t[i]) {
			ans++
		}
	}
	fmt.Println(ans)
}
