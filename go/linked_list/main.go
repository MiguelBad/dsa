package main

import "fmt"

type Value interface {
	int | string
}

type Node[T Value] struct {
	item T
	next *Node[T]
}

type LinkedList[T Value] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (ll *LinkedList[T]) insert(item T, idx int) {
	new_node := &Node[T]{item: item, next: nil}

	if ll.length == 0 {
		ll.head = new_node
		ll.tail = new_node
	} else {
		curr := ll.head
		currIdx := 0
		for curr != nil {
			if currIdx == idx {
				new_node.next = curr.next
				curr.next = new_node
			}
			currIdx++
			curr = curr.next
		}
	}

	ll.length++
}

func (ll *LinkedList[T]) insertHead(item T) {
	new_node := &Node[T]{item: item, next: nil}

	if ll.length == 0 {
		ll.head = new_node
		ll.tail = new_node
	} else {
		temp := ll.head
		ll.head = new_node
		ll.head.next = temp
	}

	ll.length++
}

func (ll *LinkedList[T]) insertTail(item T) {
	new_node := &Node[T]{item: item, next: nil}

	if ll.length == 0 {
		ll.head = new_node
		ll.tail = new_node
	} else {
		ll.tail.next = new_node
		ll.tail = new_node
	}

	ll.length++
}

func (ll *LinkedList[T]) delete(item T) *T {
	if ll.length == 0 {
		fmt.Println("List is empty")
		return nil
	}

	curr := ll.head
	for curr != nil {
		if curr.next.item == item {
			toDelete := curr.next
			curr.next = toDelete.next

			ll.length--
			return &toDelete.item
		}
		curr = curr.next
	}

	return nil
}

func (ll *LinkedList[T]) deleteHead() *T {
	if ll.length == 0 {
		fmt.Println("List is empty")
		return nil
	}

	ll.head = ll.head.next
	ll.length--
	return nil
}

func (ll *LinkedList[T]) displayList() {
	curr := ll.head

	for curr != nil {
		fmt.Print(curr.item, " - ")
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	intList := &LinkedList[int]{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	for _, i := range []int{2, 7, 4, 5} {
		intList.insertTail(i)
		intList.displayList()
	}

	for _, i := range []int{3, 1, 6, 2} {
		intList.insertHead(i)
		intList.displayList()
	}

	intList.insert(10, 5)
	intList.displayList()

	for range [4]int{} {
		intList.deleteHead()
		intList.displayList()
	}

	intList.delete(10)
    intList.displayList()

	// stringList := &LinkedList[string]{
	// 	head:   nil,
	// 	tail:   nil,
	// 	length: 0,
	// }
}
