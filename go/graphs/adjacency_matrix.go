package graphs

type AdjMatrix struct {
	matrix   [][]int
	vertices []int
}

func NewAdjMatrix() *AdjMatrix {
	adjMatrix := &AdjMatrix{
		vertices: []int{1, 2, 3, 4, 5, 6},
		matrix: [][]int{
			{0, 0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0, 1},
			{0, 1, 0, 1, 0, 0},
			{1, 0, 1, 0, 1, 0},
			{1, 0, 0, 1, 0, 1},
			{1, 0, 0, 0, 1, 0},
		},
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
	//     6

	return adjMatrix
}
