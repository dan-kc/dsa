package structs

import "fmt"

type QNode[T any] struct {
	data T
	prev *QNode[T]
}

type Queue[T any] struct {
	head   *QNode[T]
	length int
}

func (q *Queue[T]) deque() (T, error) {
	if q.length == 0 {
		return *new(T), fmt.Errorf("already empty")
	}
	oldHead := q.head
	q.head = q.head.prev
	q.length--
	return oldHead.data, nil
}

func (q *Queue[T]) enqueue(data T) {
	node := QNode[T]{data: data, prev: q.head}
	q.head = &node
	q.length++
}

func (q Queue[T]) peek() (T, error) {
	if q.head == nil {
		return *new(T), fmt.Errorf("")
	}
	return q.head.data, nil
}
