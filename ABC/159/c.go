package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)

	sq := float64(l) / 3.0
	fmt.Println(sq * sq * sq)
}
