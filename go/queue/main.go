package main

import "fmt"

type Node struct {
	Item int
	Next *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enqueue(item int) {
	newNode := &Node{Item: item}
	q.Length++

	if q.Length == 1 {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

func (q *Queue) Deque() int {
	if q.Length == 0 {
		fmt.Println("Queue is empty!")
		return 0
	}

	q.Length--
	item := q.Head.Item

	if q.Length == 0 {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}

	return item
}

func (q *Queue) Peek() int {
	if q.Length == 0 {
		fmt.Println("Queue is empty!")
		return 0
	}

	return q.Head.Item
}

func seeVal(q Queue) {
	fmt.Println("Length:", q.Length)
	fmt.Println("Head Item", q.Head.Item)
	fmt.Println("Peek:", q.Peek())
	fmt.Println("Tail Item", q.Tail.Item)
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
	fmt.Println(q.Length)
}
