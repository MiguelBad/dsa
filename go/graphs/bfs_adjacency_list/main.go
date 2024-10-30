package main

import (
	"dsa/dsa/graphs"
	"fmt"
)

type QNode struct {
	item graphs.Vertex
	next *QNode
}

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *Queue) enqueue(item graphs.Vertex) {
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

func (q *Queue) deque() graphs.Vertex {
	if q.length < 1 {
		panic("Deque-ed empty queue")
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

func findNeedle(graph graphs.AdjList, root graphs.Vertex, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}

	q := &Queue{}
	q.enqueue(root)
	seen[root.Item] = true
	for q.length > 0 {
		curr := q.deque()

		// equivalent to:
		// if curr.Item == needle { break } but less iteration
		if prev[needle] != -1 {
			break
		}

		for _, i := range curr.Edge {
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr.Item
			q.enqueue(*graph[i])
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
	adjList := graphs.NewAdjList()
	root := *adjList[0]
	needle := 0
	path := findNeedle(adjList, root, needle)

	if len(path) > 0 {
		fmt.Printf("%d is found on path:\n%v\n", needle, path)
	} else {
		fmt.Printf("%d not found\n", needle)
	}
}
