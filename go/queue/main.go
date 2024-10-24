package main

import "fmt"

type Value interface {
	string | int
}

type Node[T Value] struct {
	item T
	next *Node[T]
}

type Queue[T Value] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (q *Queue[T]) enqueue(item T) {
	newNode := &Node[T]{
		item: item,
		next: nil,
	}

	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.length++
}

func (q *Queue[T]) deque() *T {
	if q.length == 0 {
		return nil
	}

	item := q.head.item
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	return &item
}

func (q *Queue[T]) peek() *T {
	if q.length == 0 {
		return nil
	}

	return &q.head.item
}

func (q *Queue[T]) displayQueue() {
	curr := q.head
	for curr != nil {
		fmt.Print(curr.item, " - ")
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	queue := &Queue[int]{}

	for i := 0; i < 5; i++ {
		queue.enqueue(i)
		queue.displayQueue()
	}

	currHead := *queue.peek()
	fmt.Println(currHead)

	for i := 0; i < 5; i++ {
		queue.deque()
		queue.displayQueue()
	}
}
