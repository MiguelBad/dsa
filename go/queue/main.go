package main

type IntOrString interface {
	int | string
}

type Node[T IntOrString] struct {
	item T
	next *Node[T]
	prev *Node[T]
}

type Queue[T IntOrString] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (q *Queue[T]) enqueue(item T) {
	newNode := Node[T]{item: item}

	if q.length == 0 {
		q.head = &newNode
		q.tail = &newNode

	}
}

func (q *Queue[T]) deque() {
}

func (q *Queue[T]) is_empty() bool {
	return q.length == 0
}

func (q *Queue[T]) peek() T {
	return q.head.item
}

func main() {
}
