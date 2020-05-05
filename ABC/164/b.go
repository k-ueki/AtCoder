package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	x := a / d
	y := c / b

	if x >= y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
