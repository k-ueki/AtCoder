package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var a, b int
	to := make(map[int][]int)
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		a--
		b--
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
	}

	var q Queue
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
			if !seen[val] {
				pre[val] = v.(int)
				q.Enqueue(val)
			}
			seen[val] = true
		}
	}

	fmt.Println("Yes")
	for i := 1; i < n; i++ {
		fmt.Println(pre[i] + 1)
	}
}

type (
	Queue struct {
		start, end *node
		length     int
	}

	node struct {
		value interface{}
		next  *node
	}
)

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Enqueue(v interface{}) {
	n := &node{v, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return n.value
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}

func (q *Queue) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}
