package main

import "fmt"

func main() {
	test := PriorityQueue{}
	test.Push(5)
	test.Push(1)
	test.Push(9)
	test.Push(10)
	test.Push(20)
	test.Push(0)
	test.Push(22)

	test.Pop()
	test.Pop()
	fmt.Print(test.Peek())
}

type PriorityQueue struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func (q *PriorityQueue) Push(newData int) {
	// if the list is empty, then just add it
	if q.head == nil {
		q.head = &Node{data: newData, next: nil}
		return
	}

	// if the head of the list is less priority than what we're adding,
	// then add to front
	if q.head.data < newData {
		q.head = &Node{data: newData, next: q.head}
		return
	}

	if q.head.next == nil {
		q.head.next = &Node{data: newData, next: nil}
	}
	cursor := q.head
	for cursor.next.data >= newData {
		cursor = cursor.next
		if cursor.next == nil { // cursor is pointing to the last node
			break
		}
	}

	node := Node{data: newData, next: cursor.next}
	cursor.next = &node
}

func (q *PriorityQueue) Pop() error {
	if q.head == nil {
		return fmt.Errorf("empty sorry m8")
	}
	q.head = q.head.next
	return nil
}

func (q *PriorityQueue) Peek() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("empty sorry m8")
	}
	return q.head.data, nil
}
