package main

import "fmt"

// "fmt"

type BSTNode struct {
	item  *int
	left  *BSTNode
	right *BSTNode
}

type BinarySearchTree struct {
	root *BSTNode
}

func (BST *BinarySearchTree) insert(item int) {
	curr := BST.root
	new_node := BSTNode{item: &item}

	if curr.item == nil {
		BST.root = &new_node
		return
	}

	for curr != nil {
		if *curr.item > item {
			if curr.right == nil {
				curr.right = &new_node
				return
			}
			curr = curr.right
		} else {
			if curr.left == nil {
				curr.left = &new_node
				return
			}
			curr = curr.left
		}
	}

}

func (BST *BinarySearchTree) delete() {
}

func (BST *BinarySearchTree) search() {
}

func (BST *BinarySearchTree) inOrder() {
	curr := BST.root
	var path = []int{}

	BST._dfs(curr, path)
	fmt.Printf("In order path: %v", path)
}

func (BST *BinarySearchTree) _dfs(curr *BSTNode, path []int) {
}

func main() {
}
