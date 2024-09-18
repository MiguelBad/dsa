package create_binary_tree

type Node struct {
	item  int
	left  *Node
	right *Node
}

func CreateTree() *Node {
	node1 := &Node{item: 1}
	node2 := &Node{item: 2}
	node3 := &Node{item: 3}
	node4 := &Node{item: 4}
	node5 := &Node{item: 5}
	node6 := &Node{item: 6}
	node7 := &Node{item: 7}
	node8 := &Node{item: 8}
	node9 := &Node{item: 9}
	node10 := &Node{item: 10}
	node11 := &Node{item: 11}
	node12 := &Node{item: 12}
	node13 := &Node{item: 13}

	node1.left = node2
	node1.right = node3

	node2.left = node4
	node2.right = node5

	node3.left = node6
	node3.right = node7

	node4.left = node8
	node4.right = node9

	node5.left = node10
	node5.right = node11

	node5.left = node12
	node5.right = node13

    return node1

//          __  1  __
//         /         \
//        2           3
//      /   \       /   \
//     4     5     6     7
//    / \   / \   / \   / \
//   8   9 10 11 12 13 -   -

}
