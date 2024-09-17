package main

import "fmt"

type IntOrString interface {
	int | string
}

type Node[T IntOrString] struct {
	item T
	next *Node[T]
}

type Queue[T IntOrString] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (q *Queue[T]) enqueue(item T) {
}

func (q *Queue[T]) dequeue() {
}

func (q *Queue[T]) is_empty() {
}

func (q *Queue[T]) peek() {
}

func main() {
}
