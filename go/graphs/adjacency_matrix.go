package graphs

type AdjMatrix struct {
	Matrix   [][]int
	Vertices []int
}

func NewAdjMatrix() AdjMatrix {
	adjMatrix := AdjMatrix{
		Vertices: []int{0, 1, 2, 3, 4, 5},
		Matrix: [][]int{
			{0, 0, 1, 0, 0, 1},
			{0, 0, 0, 0, 1, 1},
			{1, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0, 1},
			{1, 1, 0, 0, 1, 0},
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
	//     0

	return adjMatrix
}
