package main

import (
	"dsa/dsa/createBT"
	"fmt"
)

func compare(tree1 *createBT.Node, tree2 *createBT.Node) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	if tree1.Item != tree2.Item {
		return false
	}

	return compare(tree1.Left, tree2.Left) && compare(tree1.Right, tree2.Right)
}

func main() {
	var tree1 *createBT.Node
	var tree2 *createBT.Node
	var result bool

	// should output true
	tree1 = createBT.CreateTree()
	tree2 = createBT.CreateTree()
	result = compare(tree1, tree2)
	fmt.Println(result)

	// should output false
	tree1 = createBT.CreateTree()
	tree2 = createBT.CreateTree2()
	result = compare(tree1, tree2)
	fmt.Println(result)
}
