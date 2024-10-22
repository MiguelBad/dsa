package main

import (
	"fmt"
)

type BSTNode struct {
	item  int
	left  *BSTNode
	right *BSTNode
}

func createBSTNode(item int) *BSTNode {
	return &BSTNode{
		item:  item,
		left:  nil,
		right: nil,
	}
}

type BST struct {
	root *BSTNode
}

func (b *BST) insert(item int) {
	new_node := createBSTNode(item)

	if b.root == nil {
		b.root = new_node
		return
	}

	curr := b.root
	for curr != nil {
		if curr.item > item {
			if curr.left == nil {
				curr.left = new_node
				return
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = new_node
				return
			}
			curr = curr.right
		}
	}
}

func (b *BST) delete() {
}

func (b *BST) search(item int) bool {
	curr := b.root
	if curr.item == item {
		return true
	}

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

func (b *BST) inOrder() {
	path := []int{}
	_bfs(b.root, &path)
	fmt.Println(path)
}

func _bfs(curr *BSTNode, path *[]int) {
	if curr == nil {
		return
	}

	_bfs(curr.left, path)
	*path = append(*path, curr.item)
	_bfs(curr.right, path)
}

func main() {
	var bst = &BST{}
	for _, i := range []int{3, 5, 1, 10, 5, 2} {
		bst.insert(i)
	}

	bst.inOrder()
	fmt.Println(bst.search(9))
}
