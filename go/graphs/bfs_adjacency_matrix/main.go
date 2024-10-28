package main

import (
	"dsa/dsa/graphs"
	"fmt"
)

type QNode struct {
	item int
	next *QNode
}

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *Queue) enqueue(item int) {
	newNode := &QNode{item: item}
	if q.length < 1 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.length++
}

func (q *Queue) deque() int {
	if q.length < 1 {
		panic("Deque-ed an empty queue")
	}

	item := q.head.item
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.length--
	return item
}

func findNeedle(graph graphs.AdjMatrix, needle int, root int) []int {
	seen := make([]bool, len(graph.Vertices))
	prev := make([]int, len(graph.Vertices))
	for i := 0; i < len(graph.Vertices); i++ {
		prev[i] = -1
	}

	q := &Queue{}
	q.enqueue(root)
	seen[root] = true
	for q.length > 0 {
		curr := q.deque()
		if seen[needle] {
			break
		}

		for i := 0; i < len(graph.Vertices); i++ {
			if graph.Matrix[curr][i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.enqueue(i)
		}

	}

	curr := needle
	path := []int{}
	for curr != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}
	return path
}

func main() {
	adjMatrix := graphs.NewAdjMatrix()
	needle := 0
	root := 1
	path := findNeedle(adjMatrix, needle, root)

	if len(path) > 0 {
		fmt.Printf("%d is found:\n%v\n", needle, path)
	} else {
		fmt.Printf("%d is not found\n", needle)
	}
}
