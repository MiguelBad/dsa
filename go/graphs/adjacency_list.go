package graphs

type Vertex struct {
	item int
	edge []int
}

type AdjList struct {
	vertices []*Vertex
}

func NewAdjList() *AdjList {
	adjList := &AdjList{}

	adjList.vertices = append(adjList.vertices, &Vertex{item: 1, edge: []int{4, 5}})
	adjList.vertices = append(adjList.vertices, &Vertex{item: 2, edge: []int{3, 6}})
	adjList.vertices = append(adjList.vertices, &Vertex{item: 3, edge: []int{2, 4}})
	adjList.vertices = append(adjList.vertices, &Vertex{item: 4, edge: []int{1, 3, 5}})
	adjList.vertices = append(adjList.vertices, &Vertex{item: 5, edge: []int{1, 4, 6}})
	adjList.vertices = append(adjList.vertices, &Vertex{item: 6, edge: []int{2, 5}})

	//        1
	//        | \
	//        |  \
	//        |   \
	//  3-----|----4
	//  |     |   /
	//  |     |  /
	//  |     | /
	//  |     |/
	//  2     5
	//   \   /
	//    \ /
	//     6

	return adjList
}
