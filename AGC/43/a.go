package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([][]string, h+1)
	for i := 1; i <= h; i++ {
		s[i] = make([]string, w+1)
		for j := 1; j <= w; j++ {
			fmt.Scan(&s[i][j])
		}
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
