package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DQ struct {
	head *Node
	tail *Node
	size int
}

func (q *DQ) PopFront() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("list empty")
	}

	q.size--

	if q.size == 1 {
		r := q.head.data
		q.head = nil
		q.tail = nil
		return r, nil
	}

	r := q.head.data
	q.head = q.head.next
	q.head.prev = nil // otherwise it will still point to the removed node
	return r, nil
}

func (q *DQ) PopRear() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("list empty")
	}

	q.size--

	if q.size == 1 {
		r := q.tail.data
		q.head = nil
		q.tail = nil
		return r, nil
	}

	r := q.tail.data
	q.tail = q.tail.prev
	q.tail.next = nil // otherwise it will still point to the removed node
	return r, nil
}

func (q *DQ) PushFront(data int) {
	q.size++
	if q.size == 0{
		newNode := Node{data: data, next: nil, prev: nil}
		q.head = &newNode
		q.tail = &newNode
		return
	}

	newNode := Node{data: data, next: q.head, prev: nil}
	q.head = &newNode
}

func (q *DQ) PushRear(data int) {
	q.size++
	if q.size == 0 {
		newNode := Node{data: data, next: nil, prev: nil}
		q.head = &newNode
		q.tail = &newNode
		return
	}

	newNode := Node{data: data, next: nil, prev: q.tail}
	q.tail = &newNode
}

func (q *DQ) PeekFront() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("q is empty m8")
	}
	return q.head.data, nil
}

func (q *DQ) PeekRear() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("q is empty m8")
	}
	return q.tail.data, nil
}

func (q DQ) Size() int {
	return q.size
}

func (q DQ) IsEmpty() bool {
	return q.size == 0
}
