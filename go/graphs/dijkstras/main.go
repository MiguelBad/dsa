package main

import (
	"dsa/dsa/graphs"
	"fmt"
	"math"
)

func dijkstra(graph graphs.AdjList, source graphs.Vertex, end int) []int {
	prev := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}
	distances := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		distances[i] = math.MaxInt
	}

	heap := &MinHeap{}
	heap.insert(&source, 0)

	for {
		break
	}

	path := []int{}
	return path
}

func main() {
	adjList := graphs.NewAdjList()
	source := adjList[0]
	end := 1

	path := dijkstra(adjList, *source, end)
	if len(path) > 0 {
		fmt.Printf("Shortest path to %d:\n%v\n", end, path)
	} else {
		fmt.Printf("No path leads to %d\n", end)
	}
}
