package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func NewxtSlice(num int) []int {

	rep := make([]int, num)
	in.Scan()
	ele := strings.Split(in.Text(), " ")
	for i := 0; i < num; i++ {
		n, _ := strconv.Atoi(ele[i])
		rep[i] = n
	}
	return rep
}
func iszero(num []int) bool {
	for _, v := range num {
		if v == 0 {
			return true
		}
	}
	return false
}
func max(num []int) int {
	max := 0
	for _, v := range num {
		if max < v {
			max = v
		}
	}
	return max
}
func isOrder(num []int) bool {
	for i := 1; i < len(num)-1; i++ {
		if num[i] > num[i-1] && num[i] < num[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	rep := NewxtSlice(N)
	temp := 0
	count := 0
	var Maxs []int
LABEL1:
	if iszero(rep) {
		for i := 0; i < N; i++ {
		LABEL2:
			fmt.Println("temp,i", temp, i)
			if rep[i] == 0 {
				if i == 0 {
					temp = i + 1
				} else {
					fmt.Println(temp, i)
					if temp == i {
						fmt.Println("フザケンナ")
						temp++
						i++
						goto LABEL2
					} else if isOrder(rep[temp:i]) {

						fmt.Println("A", temp, i)
						MaxNum := max(rep[temp:i])
						fmt.Println(MaxNum)
						Maxs = append(Maxs, MaxNum)
						temp = i + 1
						fmt.Println("MAXS", Maxs)
					} else {
						fmt.Println("B", temp, i)
						fmt.Println("順序が良くない")
						fmt.Println(rep[temp:i])
						for iszero(rep[temp:i]) != true {
							for k := temp; k < i; k++ {
								rep[k]--
							}
							count++
						}
						fmt.Println("countA", count)
						goto LABEL2
					}
				}
			} else if i == N-1 {
				fmt.Println("HEHEHE")
				fmt.Println(rep[temp:])
				if isOrder(rep[temp:]) {
					fmt.Println("JASHDJSAHJ")
					MaxNum := max(rep[temp:])
					Maxs = append(Maxs, MaxNum)
					fmt.Println(MaxNum)
				} else {
					fmt.Println("順番が良くない")
					fmt.Println(temp, i)
					fmt.Println(rep[temp : i+1])
					for iszero(rep[temp:i+1]) != true {
						for k := temp; k <= i; k++ {
							rep[k]--
						}
						count++
					}
					temptemp := temp
				LABEL3:
					if iszero(rep[temp : i+1]) {
						for t, v := range rep[temp : i+1] {
							if v == 0 {
								fmt.Println("TEMP", temp, "T", t)
								fmt.Println(max(rep[temp : temp+t]))
								Maxs = append(Maxs, max(rep[temp:temp+t]))
								temp = temp + t + 1
								goto LABEL3
							}
						}
					}
					temp = temptemp
					fmt.Println("countB", count)
					temp++
					fmt.Println("TEMP,I", temp, i)
					fmt.Println("REP", rep)

					goto LABEL2
				}
			}

		}
	} else {
		for iszero(rep) != true {
			for i := 0; i < len(rep); i++ {
				rep[i]--
			}
			count++
			fmt.Println("countC", count)
			fmt.Println("rep", rep)
			goto LABEL1
		}
	}

	ans := 0
	for _, v := range Maxs {
		ans += v
	}
	ans += count
	fmt.Println("count", count)
	fmt.Println(Maxs)
	fmt.Println(ans)
}
