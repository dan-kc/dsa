package structs

import "fmt"

type llNode[T any] struct {
	data T
	next *llNode[T]
}

type LinkedList[T any] struct {
	head *llNode[T]
}

func (l LinkedList[T]) Length() int {
	len := 0
	cursor := l.head
	for cursor != nil {
		len++
		cursor = cursor.next
	}
	return len
}

func (l *LinkedList[T]) InsertAt(item T, index int) error {
	cursor := l.head
	cursorIdx := 0
	if index == 0 {
		l.Prepend(item)
	}
	for cursorIdx != index-1 {
		if cursor == nil {
			return fmt.Errorf("Soz")
		}
		cursor = cursor.next
		cursorIdx++
	}
	node := llNode[T]{data: item, next: cursor.next}
	cursor.next = &node

	return nil
}

func (l *LinkedList[T]) Prepend(item T) {
	node := llNode[T]{data: item, next: l.head}
	l.head = &node
}

func (l *LinkedList[T]) Append(item T) {
	node := llNode[T]{data: item, next: nil}
	if l.head == nil {
		l.head = &node
	}
	cursor := l.head
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = &node
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	cursor := l.head
	for i := 1; i < index+1; i++ {
		if cursor.next == nil {
			return *new(T), fmt.Errorf("index value empty")
		}
		cursor = cursor.next
	}
	if cursor == nil {
		return *new(T), fmt.Errorf("index value empty")
	}
	return cursor.data, nil
}

func (l *LinkedList[T]) RemoveAt(index int) (T, error) {
	cursor := l.head
	for i := 1; i < index; i++ {
		cursor = cursor.next
	}
	removedData := cursor.next.data
	cursor.next = cursor.next.next
	return removedData, nil
}
