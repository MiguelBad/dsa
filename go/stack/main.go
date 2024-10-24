package main

import "fmt"

type Value interface {
	string | int
}

type Node[T Value] struct {
	item T
	next *Node[T]
}

type Stack[T Value] struct {
	head   *Node[T]
	length int
}

func (s *Stack[T]) push(item T) {
	newNode := &Node[T]{
		item: item,
		next: nil,
	}

	if s.length == 0 {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}

	s.length++
}

func (s *Stack[T]) pop() *T {
	if s.length == 0 {
		return nil
	}

	item := s.head.item
	if s.length == 1 {
		s.head = nil
	} else {
		s.head = s.head.next
	}

	s.length--
	return &item
}

func (s *Stack[T]) peek() T {
	return s.head.item
}

func (s *Stack[T]) displayStack() {
	curr := s.head
	for curr != nil {
		fmt.Print(curr.item, " - ")
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	stack := &Stack[int]{
		head:   nil,
		length: 0,
	}

	for i := 0; i < 5; i++ {
		stack.push(i)
		stack.displayStack()
	}

	currHead := stack.peek()
	fmt.Println(currHead)

	for i := 0; i < 5; i++ {
		stack.pop()
		stack.displayStack()
	}
}
