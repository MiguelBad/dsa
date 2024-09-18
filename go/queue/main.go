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

	if q.length == 1 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
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

func (q *Queue) Peek() int {
	if q.length == 0 {
		fmt.Println("Queue is empty!")
		return 0
	}

	return q.head.item
}

func seeVal(q Queue) {
	fmt.Println("Length:", q.length)
	fmt.Println("Head Item", q.head.item)
	fmt.Println("Peek:", q.Peek())
	fmt.Println("Tail Item", q.tail.item)
	fmt.Println()
}

func main() {
	q := &Queue{}

	q.Enqueue(4)
	seeVal(*q)
	q.Enqueue(2)
	seeVal(*q)
	q.Enqueue(9)
	seeVal(*q)
	q.Enqueue(10)
	seeVal(*q)

	q.Deque()
	seeVal(*q)
	q.Deque()
	seeVal(*q)
	q.Deque()
	seeVal(*q)
	q.Deque()
    fmt.Println(q.length)
}
