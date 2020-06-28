package main

import (
	"github.com/k-ueki/AtCoder/libs/data"
)

func main() {
	var q data.Queue
	n := 10
	seen := make(map[int]bool, n)
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = -1
		seen[i] = false
	}
	q.Enqueue(0)
	seen[0] = true
	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, val := range to[v.(int)] {
			// to := [int][]int
			if !seen[val] {
				pre[val] = v.(int)
				q.Enqueue(val)
			}
			seen[val] = true
		}
	}
}
