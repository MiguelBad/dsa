package main

import "fmt"

type Node struct {
	item int
	next *Node
}

type Stack struct {
	head   *Node
	length int
}

func (s *Stack) Push(item int) {
	newNode := &Node{item: item}
	s.length++

	if s.length == 1 {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		fmt.Println("Stack is empty")
		return 0
	}

	item := s.head.item
	s.length--

	if s.length == 0 {
		s.head = nil
	} else {
		s.head = s.head.next
	}

	return item
}

func (s *Stack) Peek() int {
	if s.length == 0 {
		fmt.Println("Stack is empty")
		return 0
	}

	return s.head.item
}

func main() {
	s := &Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Peek())

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
