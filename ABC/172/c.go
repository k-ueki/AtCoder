package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	q1 := NewQueue()
	q2 := NewQueue()
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		q1.Enqueue(tmp)
	}
	for i := 0; i < m; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		q2.Enqueue(tmp)
	}

	sum := 0
	ans := 0
	for sum <= k {
		if !q1.IsEmpty() && !q2.IsEmpty() {
			if q1.Peek().(int) > q2.Peek().(int) {
				sum += q2.Dequeue().(int)
				ans++
			} else {
				sum += q1.Dequeue().(int)
				ans++
			}
			continue
		}
		if q1.IsEmpty() && q2.IsEmpty() {
			fmt.Println(ans)
			return
		}
		if q1.IsEmpty() {
			sum += q2.Dequeue().(int)
			ans++
			continue
		}
		if q2.IsEmpty() {
			sum += q1.Dequeue().(int)
			ans++
			continue
		}
	}
	fmt.Println(ans - 1)
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

func NewQueue() *Queue {
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
