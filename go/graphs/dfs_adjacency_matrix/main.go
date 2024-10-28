package main

import (
	"dsa/dsa/graphs"
	"fmt"
)

func dfs(adjMatrix graphs.AdjMatrix, needle int, curr int, path *[]int, seen []bool) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true
	*path = append(*path, curr)

	if curr == needle {
		return true
	}

	for i := 0; i < len(adjMatrix.Vertices); i++ {
		if adjMatrix.Matrix[curr][i] == 1 {
			if dfs(adjMatrix, needle, i, path, seen) {
				return true
			}
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func findNeedle(adjMatrix graphs.AdjMatrix, needle int, root int) (bool, []int) {
	path := []int{}
	seen := make([]bool, len(adjMatrix.Vertices))
	found := dfs(adjMatrix, needle, root, &path, seen)
	return found, path
}

func main() {
	adjMatrix := graphs.NewAdjMatrix()
	needle := 0

	found, path := findNeedle(adjMatrix, needle, 1)

	if found {
		fmt.Printf("%d found on path:\n%v\n", needle, path)
	} else {
		fmt.Printf("%d not found\n", needle)
	}
}
