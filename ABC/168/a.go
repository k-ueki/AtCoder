package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	n %= 10
	if n == 3 {
		fmt.Println("bon")
	} else if n == 0 || n == 1 || n == 6 || n == 8 {
		fmt.Println("pon")
	} else {
		fmt.Println("hon")
	}
}
