package main

import "fmt"

type Node struct {
	item int
	next *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) Enqueue(item int) {
	newNode := &Node{item: item}
	q.length++

	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Deque() int {
	if q.length == 0 {
		fmt.Println("Queue is empty!")
		return 0
	}

	q.length--
	item := q.head.item

	if q.length == 0 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	return item
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Peek() int {
	if q.length == 0 {
		fmt.Println("Queue is empty!")
		return 0
	}

	return q.head.item
}

func main() {
}
