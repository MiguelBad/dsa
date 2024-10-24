package main

import (
	"fmt"
)

type AVLNode struct {
	item  int
	left  *AVLNode
	right *AVLNode
}

type AVL struct {
	root *AVLNode
}

type QNode struct {
	item *AVLNode
	next *QNode
}

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *Queue) enqueue(item *AVLNode) {
	newNode := &QNode{item: item}
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.length++
}

func (q *Queue) deque() *AVLNode {
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

	q.length--
	return item
}

func (bst *AVL) insert(item int) {
	newNode := &AVLNode{item: item}
	if bst.root == nil {
		bst.root = newNode
		return
	}

	curr := bst.root
	for curr != nil {
		if item < curr.item {
			if curr.left == nil {
				curr.left = newNode
				return
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = newNode
				return
			}
			curr = curr.right
		}
	}
}

func (bst *AVL) delete(item int) *int {
	if bst.root == nil {
		return nil
	}
	bst.root = bst._deleteHelper(bst.root, item)
	return &item
}

func (bst *AVL) _deleteHelper(curr *AVLNode, item int) *AVLNode {
	if curr == nil {
		return nil
	} else if curr.item > item {
		curr.left = bst._deleteHelper(curr.left, item)
	} else if curr.item < item {
		curr.right = bst._deleteHelper(curr.right, item)
	} else {
		if curr.left == nil {
			curr = curr.right
		} else if curr.right == nil {
			curr = curr.left
		} else {
			successor := bst.findSuccessor(curr.right)
			curr.item = successor.item
			curr.right = bst._deleteHelper(curr.right, successor.item)
		}
	}
	return curr
}

func (bst *AVL) findSuccessor(node *AVLNode) *AVLNode {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}

	return curr
}

func (bst *AVL) search(item int) bool {
	if bst.root == nil {
		return false
	}

	curr := bst.root
	for curr != nil {
		if curr.item == item {
			return true
		}
		if curr.item > item {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (bst *AVL) rotate() {
}

func (bst *AVL) dfs() {
	path := []int{}
	dfsHelper(bst.root, &path)
	fmt.Println(path)
}

func dfsHelper(curr *AVLNode, path *[]int) {
	if curr == nil {
		return
	}

	dfsHelper(curr.left, path)
	*path = append(*path, curr.item)
	dfsHelper(curr.right, path)
}

func (bst *AVL) bfs() {
	path := []int{}
	queue := &Queue{}
	queue.enqueue(bst.root)

	for queue.length != 0 {
		item := queue.deque()
		if item == nil {
			continue
		}

		path = append(path, item.item)
		queue.enqueue(item.left)
		queue.enqueue(item.right)
	}

	fmt.Println(path)
}

func main() {
	bst := &AVL{}
	bst.insert(3)
	bst.bfs()
	bst.dfs()
	bst.insert(2)
	bst.bfs()
	bst.dfs()
	bst.insert(9)
	bst.bfs()
	bst.dfs()
	bst.insert(7)
	bst.bfs()
	bst.dfs()

    bst.delete(3)
    bst.bfs()
    bst.dfs()
    bst.delete(7)
    bst.bfs()
    bst.dfs()
    bst.delete(2)
    bst.bfs()
    bst.dfs()
    bst.delete(9)
    bst.bfs()
    bst.dfs()
}
