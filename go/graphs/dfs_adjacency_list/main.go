package main

import (
	"dsa/dsa/graphs"
	"fmt"
)

func dfs(adjList graphs.AdjList, needle int, curr graphs.Vertex, path *[]int, seen []bool) bool {
	if seen[curr.Item] {
		return false
	}

	seen[curr.Item] = true
	*path = append(*path, curr.Item)

	if curr.Item == needle {
		return true
	}

	for _, i := range curr.Edge {
		if dfs(adjList, needle, *adjList[i], path, seen) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func findNeedle(adjList graphs.AdjList, needle int, root graphs.Vertex) (bool, []int) {
	path := []int{}
	seen := make([]bool, len(adjList))
	found := dfs(adjList, needle, root, &path, seen)
	return found, path
}

func main() {
	adjList, root := graphs.NewAdjList()
	needle := 3

	found, path := findNeedle(adjList, needle, root)

	if found {
		fmt.Printf("%d found on path:\n%v\n", needle, path)
	} else {
		fmt.Printf("Needle not found")
	}
}
