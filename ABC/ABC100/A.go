package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)

	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
		return
	}
	fmt.Println(":(")
}
