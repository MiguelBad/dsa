package main

import (
	"fmt"
	"math"
)

func dijkstra(graph *AdjList, source *Vertex, end int) []int {
	prev := make([]int, len(*graph))
	for i := 0; i < len(*graph); i++ {
		prev[i] = -1
	}
	distances := make([]int, len(*graph))
	for i := 0; i < len(*graph); i++ {
		distances[i] = math.MaxInt
	}

	distances[source.item] = 0
	heap := &MinHeap{}
	heap.insert(source, 0)

	for len(heap.elements) > 0 {
		curr := heap.delete()

		for _, edge := range curr.edge {
			dist := distances[curr.item] + edge.weight

			if dist < distances[edge.to] {
				distances[edge.to] = dist
				prev[edge.to] = curr.item
				heap.insert((*graph)[edge.to], edge.weight)
			}
		}
	}

	path := []int{}
	curr := end
	for prev[curr] != -1 {
		path = append(path, prev[curr])
		curr = prev[curr]
	}

	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	if len(path) > 0 {
		path = append(path, end)
	}

	return path
}

func main() {
	adjList := NewAdjList()
	source := adjList[0]
	end := 1

	path := dijkstra(&adjList, source, end)
	if len(path) > 0 {
		fmt.Printf("Shortest path to %d:\n%v\n", end, path)
	} else {
		fmt.Printf("No path leads to %d\n", end)
	}
}
