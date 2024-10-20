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

func newBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
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
	BST._dfs(curr.left, path)
	path = append(path, *curr.item)
	BST._dfs(curr.right, path)
}

func main() {
	var bst = newBinarySearchTree()

	items := []int{5, 1, 6, 2, 3}
	for i := range items {
		bst.insert(i)
	}

	bst.inOrder()
}
