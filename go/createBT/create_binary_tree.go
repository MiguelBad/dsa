package createBT

type Node struct {
	Item  int
	Left  *Node
	Right *Node
}

func CreateTree() *Node {
	node1 := &Node{Item: 1}
	node2 := &Node{Item: 2}
	node3 := &Node{Item: 3}
	node4 := &Node{Item: 4}
	node5 := &Node{Item: 5}
	node6 := &Node{Item: 6}
	node7 := &Node{Item: 7}
	node8 := &Node{Item: 8}
	node9 := &Node{Item: 9}
	node10 := &Node{Item: 10}
	node11 := &Node{Item: 11}
	node12 := &Node{Item: 12}
	node13 := &Node{Item: 13}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node4.Left = node8
	node4.Right = node9

	node5.Left = node10
	node5.Right = node11

	node6.Left = node12
	node6.Right = node13

	return node1

	//          __  1  __
	//         /         \
	//        2           3
	//      /   \       /   \
	//     4     5     6     7
	//    / \   / \   / \   / \
	//   8   9 10 11 12 13 -   -

}

func CreateTree2() *Node {
	node1 := &Node{Item: 1}
	node2 := &Node{Item: 2}
	node3 := &Node{Item: 3}
	node4 := &Node{Item: 4}
	node5 := &Node{Item: 5}
	node6 := &Node{Item: 6}
	node7 := &Node{Item: 7}
	node8 := &Node{Item: 8}
	node9 := &Node{Item: 9}
	node10 := &Node{Item: 10}
	node11 := &Node{Item: 11}
	node12 := &Node{Item: 12}
	node13 := &Node{Item: 20}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node4.Left = node8
	node4.Right = node9

	node5.Left = node10
	node5.Right = node11

	node6.Left = node12
	node6.Right = node13

	return node1

	//          __  1  __
	//         /         \
	//        2           3
	//      /   \       /   \
	//     4     5     6     7
	//    / \   / \   / \   / \
	//   8   9 10 11 12 20 -   -

}
