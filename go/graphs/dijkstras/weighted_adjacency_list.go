package main

type GraphEdge struct {
	to     int
	weight int
}

type Vertex struct {
	item int
	edge []*GraphEdge
}

type AdjList []*Vertex

func NewAdjList() AdjList {
	_4_1 := &GraphEdge{to: 1, weight: 2}
	_1_4 := &GraphEdge{to: 4, weight: 2}
	_5_1 := &GraphEdge{to: 1, weight: 6}
	_1_5 := &GraphEdge{to: 5, weight: 6}
	_4_3 := &GraphEdge{to: 3, weight: 5}
	_3_4 := &GraphEdge{to: 4, weight: 5}
	_4_5 := &GraphEdge{to: 5, weight: 2}
	_5_4 := &GraphEdge{to: 4, weight: 2}
	_2_3 := &GraphEdge{to: 3, weight: 3}
	_3_2 := &GraphEdge{to: 2, weight: 3}
	_5_0 := &GraphEdge{to: 0, weight: 2}
	_0_5 := &GraphEdge{to: 5, weight: 2}
	_2_0 := &GraphEdge{to: 0, weight: 2}
	_0_2 := &GraphEdge{to: 2, weight: 2}

	var adjList AdjList = []*Vertex{
		{item: 0, edge: []*GraphEdge{_0_2, _0_5}},
		{item: 1, edge: []*GraphEdge{_1_4, _1_5}},
		{item: 2, edge: []*GraphEdge{_2_3, _2_0}},
		{item: 3, edge: []*GraphEdge{_3_2, _3_4}},
		{item: 4, edge: []*GraphEdge{_4_1, _4_3, _4_5}},
		{item: 5, edge: []*GraphEdge{_5_1, _5_4, _5_0}},
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

	return adjList
}
