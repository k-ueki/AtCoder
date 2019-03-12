package main

import "fmt"

func main() {
	var num1, num2, num3, num4 int
	var ans = "NO"
	fmt.Scanf("%d %d %d %d\n", &num1, &num2, &num3, &num4)

	if num1 == 1 || num2 == 1 || num3 == 1 || num4 == 1 {
		if num1 == 9 || num2 == 9 || num3 == 9 || num4 == 9 {
			if num1 == 7 || num2 == 7 || num3 == 7 || num4 == 7 {
				if num1 == 4 || num2 == 4 || num3 == 4 || num4 == 4 {
					ans = "YES"
				}
			}
		}
	}
	fmt.Println(ans)
}
