package main

import (
	"fmt"
)

type AVLNode struct {
	item   int
	left   *AVLNode
	right  *AVLNode
	height int
}

type AVL struct {
	root *AVLNode
}

func (avl *AVL) getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (avl *AVL) updateHeight(node *AVLNode) {
	if node == nil {
		return
	}

	leftHeight := avl.getHeight(node.left)
	rightHeight := avl.getHeight(node.right)

	node.height = 1 + max(leftHeight, rightHeight)
}

func (avl *AVL) getBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}

	leftHeight := avl.getHeight(node.left)
	rightHeight := avl.getHeight(node.right)

	return leftHeight - rightHeight
}

func (avl *AVL) leftRotate(x *AVLNode) *AVLNode {
	if x.right == nil {
		return x
	}
	y := x.right
	z := y.left

	y.left = x
	x.right = z

	avl.updateHeight(x)
	avl.updateHeight(y)

	return y
}

func (avl *AVL) rightRotate(y *AVLNode) *AVLNode {
	if y.left == nil {
		return y
	}

	x := y.left
	z := x.right

	x.right = y
	y.left = z

	avl.updateHeight(x)
	avl.updateHeight(y)

	return x
}

func newAVLNode(item int) *AVLNode {
	return &AVLNode{
		item:   item,
		height: 1,
	}
}

func (avl *AVL) insert(item int) {
	avl.root = avl.insertHelper(avl.root, item)
}

func (avl *AVL) insertHelper(curr *AVLNode, item int) *AVLNode {
	if curr == nil {
		return newAVLNode(item)
	} else if curr.item > item {
		curr.left = avl.insertHelper(curr.left, item)
	} else if curr.item < item {
		curr.right = avl.insertHelper(curr.right, item)
	} else {
		return newAVLNode(item)
	}

	avl.updateHeight(curr)
	balance := avl.getBalance(curr)

	if balance > 1 && curr.left.item > item {
		return avl.rightRotate(curr)
	}

	if balance > 1 && curr.left.item < item {
		curr.left = avl.leftRotate(curr.left)
		return avl.rightRotate(curr)
	}

	if balance < -1 && curr.right.item > item {
		curr.right = avl.rightRotate(curr.right)
		return avl.leftRotate(curr)
	}

	if balance < -1 && curr.right.item < item {
		return avl.leftRotate(curr)
	}

	return curr
}

func (avl *AVL) delete(item int) *int {
	if avl.root == nil {
		return nil
	}
	avl.root = avl.deleteHelper(avl.root, item)
	return &item
}

func (avl *AVL) deleteHelper(curr *AVLNode, item int) *AVLNode {
	if curr == nil {
		return nil
	} else if curr.item > item {
		curr.left = avl.deleteHelper(curr.left, item)
	} else if curr.item < item {
		curr.right = avl.deleteHelper(curr.right, item)
	} else {
		if curr.left == nil {
			curr = curr.right
		} else if curr.right == nil {
			curr = curr.left
		} else {
			successor := avl.findSuccessor(curr.right)
			curr.item = successor.item
			curr.right = avl.deleteHelper(curr.right, successor.item)
		}
	}

	avl.updateHeight(curr)
	balance := avl.getBalance(curr)

	if balance > 1 && avl.getBalance(curr.left) >= 0 {
		return avl.rightRotate(curr)
	}

	if balance > 1 && avl.getBalance(curr.left) <= 0 {
		curr.left = avl.leftRotate(curr.left)
		return avl.rightRotate(curr)
	}

	if balance < -1 && avl.getBalance(curr.right) >= 0 {
		curr.left = avl.rightRotate(curr.right)
		return avl.leftRotate(curr)
	}

	if balance < -1 && avl.getBalance(curr.right) <= 0 {
		return avl.leftRotate(curr)
	}

	return curr
}

func (avl *AVL) findSuccessor(node *AVLNode) *AVLNode {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}

	return curr
}

func (avl *AVL) search(item int) bool {
	if avl.root == nil {
		return false
	}

	curr := avl.root
	for curr != nil {
		if curr.item == item {
			return true
		}
		if curr.item > item {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (avl *AVL) dfs() {
	path := []int{}
	dfsHelper(avl.root, &path)
	fmt.Println(path)
}

func dfsHelper(curr *AVLNode, path *[]int) {
	if curr == nil {
		return
	}

	dfsHelper(curr.left, path)
	*path = append(*path, curr.item)
	dfsHelper(curr.right, path)
}

type QNode struct {
	item *AVLNode
	next *QNode
}

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *Queue) enqueue(item *AVLNode) {
	newNode := &QNode{item: item}
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.length++
}

func (q *Queue) deque() *AVLNode {
	if q.length == 0 {
		return nil
	}

	item := q.head.item
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	q.length--
	return item
}

func (avl *AVL) bfs() {
	path := []int{}
	queue := &Queue{}
	queue.enqueue(avl.root)

	for queue.length != 0 {
		item := queue.deque()
		if item == nil {
			continue
		}

		path = append(path, item.item)
		queue.enqueue(item.left)
		queue.enqueue(item.right)
	}

	fmt.Println(path)
}

func main() {
	avl := &AVL{}

	for _, i := range []int{3, 5, 1, 6, 9, 10} {
		avl.insert(i)
		avl.bfs()
		avl.dfs()
		fmt.Println()
	}

	for _, i := range []int{3, 5, 1, 6, 9, 10} {
		avl.delete(i)
		avl.bfs()
		avl.dfs()
		fmt.Println()
	}
}
