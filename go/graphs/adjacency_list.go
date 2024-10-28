package graphs

type Vertex struct {
	Item int
	Edge []int
}

type AdjList []*Vertex

func NewAdjList() (AdjList, Vertex) {
	var adjList AdjList = []*Vertex{
		{Item: 0, Edge: []int{2, 5}},
		{Item: 1, Edge: []int{4, 5}},
		{Item: 2, Edge: []int{3, 0}},
		{Item: 3, Edge: []int{2, 4}},
		{Item: 4, Edge: []int{1, 3, 5}},
		{Item: 5, Edge: []int{1, 4, 0}},
	}

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
	//     0

	return adjList, *adjList[1]
}
