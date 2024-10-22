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
		if curr.item == item {
			return
		}

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

func (b *BST) delete(item int) int {
	b.root = _deleteHelper(b.root, item)
	return item
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
		}

		successor := _findSuccessor(curr.right)
		curr.item = successor.item
		curr.right = _deleteHelper(curr.right, successor.item)
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
	for _, i := range []int{3, 5, 1, 10, 5, 2, 4, 1, 35, 6, 10, 3} {
		fmt.Printf("Inserting %v\n", i)
		bst.insert(i)
	}

	bst.inOrder()
	fmt.Println(bst.search(1))

	for _, i := range []int{3, 5, 1, 10, 2, 4, 35, 6} {
		fmt.Printf("Deleting %v\n", i)
		bst.delete(i)
		bst.inOrder()
	}
}
