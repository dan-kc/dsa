package structs

import "fmt"

type sNode[T any] struct {
	data T
	prev *sNode[T]
}

type Stack[T any] struct {
	head   *sNode[T]
	length int
}

func (q *Stack[T]) pop() (T, error) {
	if q.length == 0 {
		return *new(T), fmt.Errorf("already empty")
	}
	oldHead := q.head
	q.head = q.head.prev
	q.length--
	return oldHead.data, nil
}

func (q *Stack[T]) push(data T) {
	node := sNode[T]{data: data, prev: q.head}
	q.head = &node
	q.length++
}

func (q Stack[T]) peek() (T, error) {
	if q.head == nil {
		return *new(T), fmt.Errorf("")
	}
	return q.head.data, nil
}
