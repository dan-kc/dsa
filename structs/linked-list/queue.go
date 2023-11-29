package structs

import "fmt"

type qNode[T any] struct {
	data T
	next *qNode[T]
}

type Queue[T any] struct {
	length int
	head   *qNode[T]
}

func (q Queue[T]) peek() (T, error) {
	if q.length == 0 {
		return *new(T), fmt.Errorf("empty")
	}
	return q.head.data, nil
}

func (q *Queue[T]) enqueue(data T) {
	node := qNode[T]{data: data, next: nil}
	if q.length == 0 {
		q.head = &node
		q.length++
	}
	cursor := q.head
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = &node
	q.length++
}

func (q *Queue[T]) deque() (T, error) {
	if q.length == 0 {
		return *new(T), fmt.Errorf("empty")
	}
	data := q.head.data
	q.head = q.head.next
	return data, nil
}
