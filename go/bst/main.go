package main

import (
	"fmt"
)

type BSTNode struct {
	item  int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func (b *BST) insert(item int) {
	fmt.Println("Inserting:", item)
	newNode := &BSTNode{
		item:  item,
		left:  nil,
		right: nil,
	}

	curr := b.root
	if curr == nil {
		b.root = newNode
	}

	for curr != nil {
		if curr.item == item {
			fmt.Println("Item already exists")
			return
		} else if curr.item > item {
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

func (b *BST) delete(item int) {
	fmt.Println("Deleting:", item)
	b.root = _deleteHelper(b.root, item)
}

func _deleteHelper(curr *BSTNode, item int) *BSTNode {
	if curr == nil {
		return nil
	} else if curr.item > item {
		curr.left = _deleteHelper(curr.left, item)
	} else if curr.item < item {
		curr.right = _deleteHelper(curr.right, item)
	} else {
		if curr.left == nil {
			return curr.right
		} else if curr.right == nil {
			return curr.left
		} else {
			successor := _findSuccessor(curr.right)
			curr.item = successor.item
			curr.right = _deleteHelper(curr.right, successor.item)
		}
	}
	return curr
}

func _findSuccessor(curr *BSTNode) *BSTNode {
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (b *BST) search(item int) bool {
	fmt.Println("Searching", item)
	curr := b.root
	for curr != nil {
		if curr.item == item {
			return true
		} else if curr.item > item {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (b *BST) inOrder() {
	path := []int{}
	_inOrderHelper(b.root, &path)
	fmt.Println(path)
}

func _inOrderHelper(curr *BSTNode, path *[]int) {
	if curr == nil {
		return
	}

	_inOrderHelper(curr.left, path)
	*path = append(*path, curr.item)
	_inOrderHelper(curr.right, path)
}

func main() {
	bst := &BST{}

	for _, i := range []int{3, 5, 1, 10, 5, 2, 4, 1, 35, 6, 10, 3} {
		bst.insert(i)
		bst.inOrder()
	}

	fmt.Println(bst.search(1))

	for _, i := range []int{3, 5, 1, 10, 2, 4, 35, 6} {
		bst.delete(i)
		bst.inOrder()
	}
}
