package main

import (
	"dsa/dsa/create_binary_tree"
	"fmt"
)

var tree = create_binary_tree.CreateTree()

func dfsPreorder(path []int, needle int, curr *create_binary_tree.Node) ([]int, bool) {
	if curr == nil {
		return path, false
	}

	path = append(path, curr.Item)
	if curr.Item == needle {
		return path, true
	}

	if newPath, found := dfsPreorder(path, needle, curr.Left); found {
		return newPath, true
	}

	return dfsPreorder(path, needle, curr.Right)
}

func dfsInorder(path []int, needle int, curr *create_binary_tree.Node) ([]int, bool) {
	var found bool

	if curr == nil {
		return path, false
	}

	if path, found = dfsInorder(path, needle, curr.Left); found {
		return path, true
	}

	path = append(path, curr.Item)
	if curr.Item == needle {
		return path, true
	}

	return dfsInorder(path, needle, curr.Right)
}

func dfsPostorder(path []int, needle int, curr *create_binary_tree.Node) ([]int, bool) {
	var found bool

	if curr == nil {
		return path, false
	}

	if path, found = dfsPostorder(path, needle, curr.Left); found {
		return path, true
	}

	if path, found = dfsPostorder(path, needle, curr.Right); found {
		return path, true
	}

	path = append(path, curr.Item)
	if curr.Item == needle {
		return path, true
	}

    return path, false
}

func main() {
	needle := 9
	fmt.Printf("Finding number %d\n", needle)

	prePath, _ := dfsPreorder([]int{}, needle, tree)
	inPath, _ := dfsInorder([]int{}, needle, tree)
	postPath, _ := dfsPostorder([]int{}, needle, tree)

	fmt.Println("Pre-Order Path:", prePath)
	fmt.Println("In-Order Path", inPath)
	fmt.Println("Post-Order Path:", postPath)
}
