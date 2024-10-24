package main

import (
	"fmt"

	"dsa/dsa/createBT"
)

type Node struct {
	item createBT.Node
	next *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) enqueue(item createBT.Node) {
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

func (q *Queue) deque() createBT.Node {
	item := q.head.item
	q.length--

	if q.length == 0 {
		q.tail = nil
		q.head = nil
	} else {
		q.head = q.head.next
	}

	return item
}

func bfs_binary_tree(tree *createBT.Node, path []int) []int {
	q := &Queue{}
	q.enqueue(*tree)

	for q.length > 0 {
		var curr = q.deque()

		path = append(path, curr.Item)

		if curr.Left != nil {
			q.enqueue(*curr.Left)
		}

		if curr.Right != nil {
			q.enqueue(*curr.Right)
		}
	}

	return path
}

func main() {
	var tree = createBT.CreateTree()
	path := bfs_binary_tree(tree, []int{})
	fmt.Println("Tree structure:", path)
}
