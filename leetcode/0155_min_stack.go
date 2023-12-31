package main

func main() {}

type MinStack struct {
	top *Node
	min int
}

type Node struct {
	data    int
	next    *Node
	lastmin int
}

func (stack *MinStack) Push(val int) {
	var newNode *Node
	if stack.top == nil {
		newNode = &Node{data: val, next: stack.top}
		stack.min = val
	} else {
		newNode = &Node{data: val, next: stack.top, lastmin: stack.min}
	}
	stack.top = newNode
	if val < stack.min {
		stack.min = val
	}
}

func (stack *MinStack) Pop() {
	if stack.top.next == nil {
		stack.top = nil
		return
	}
	stack.min = stack.top.lastmin
	stack.top = stack.top.next
}

func (this MinStack) Top() int {
	return this.top.data
}

func (this MinStack) GetMin() int {
	return this.min
}
