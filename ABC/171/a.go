package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n < "a" {
		fmt.Println("A")
		return
	}
	fmt.Println("a")
}
