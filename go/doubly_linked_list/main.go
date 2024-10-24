package main

import "fmt"

type Value interface {
	string | int
}

type Node[T Value] struct {
	item T
	next *Node[T]
	prev *Node[T]
}

type DoublyLL[T Value] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (dll *DoublyLL[T]) insert(item T, idx int) {
	new_node := &Node[T]{
		item: item,
		next: nil,
		prev: nil,
	}

	if dll.length == 0 {
		dll.head = new_node
		dll.tail = new_node
	} else {
		curr := dll.head
		currIdx := 0
		for curr != nil {
			if currIdx == idx {
				new_node.next = curr.next
				new_node.prev = curr
				curr.next.prev = new_node
				curr.next = new_node
				break
			}

			currIdx++
			curr = curr.next
		}
	}

	dll.length++
}

func (dll *DoublyLL[T]) insertHead(item T) {
	new_node := &Node[T]{
		item: item,
		next: nil,
		prev: nil,
	}

	if dll.length == 0 {
		dll.head = new_node
		dll.tail = new_node
	} else {
		dll.head.prev = new_node
		new_node.next = dll.head
		dll.head = new_node
	}

	dll.length++
}

func (dll *DoublyLL[T]) insertTail(item T) {
	new_node := &Node[T]{
		item: item,
		next: nil,
		prev: nil,
	}

	if dll.length == 0 {
		dll.head = new_node
		dll.tail = new_node
	} else {
		dll.tail.next = new_node
		new_node.prev = dll.tail
		dll.tail = new_node
	}

	dll.length++
}

func (dll *DoublyLL[T]) delete(item T) *T {
	if dll.length < 1 {
		fmt.Println("List is empty")
		return nil
	}

	curr := dll.head
	for curr != nil {
		if curr.item == item {
			toDelete := curr

			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			curr = curr.next
			dll.length--

			return &toDelete.item
		}

		curr = curr.next
	}

	return nil
}

func (dll *DoublyLL[T]) deleteTail() *T {
	var toDelete *T

	if dll.length < 1 {
		fmt.Println("List is empty")
	} else if dll.length == 1 {
		toDelete = &dll.tail.item
		dll.head = nil
		dll.tail = nil
		dll.length--
	} else {
		toDelete = &dll.tail.item
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.length--
	}

	return toDelete
}

func (dll *DoublyLL[T]) deleteHead() *T {
	var toDelete *T

	if dll.length < 1 {
		fmt.Println("List is empty")
	} else if dll.length == 1 {
		toDelete = &dll.head.item
		dll.head = nil
		dll.tail = nil
		dll.length--
	} else {
		toDelete = &dll.head.item
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.length--
	}

	return toDelete
}

func (dll *DoublyLL[T]) search(item T) bool {
	curr := dll.head

	for curr != nil {
		if curr.item == item {
			return true
		}
		curr = curr.next
	}

	return false
}

func (dll *DoublyLL[T]) displayDLL() {
	curr := dll.head

	for curr != nil {
		fmt.Print(curr.item, " - ")
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	dll := &DoublyLL[int]{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	for _, i := range []int{2, 4, 6, 8, 10} {
		dll.insertTail(i)
		dll.displayDLL()
	}

	for _, i := range []int{1, 3, 5, 7, 9} {
		dll.insertHead(i)
		dll.displayDLL()
	}

	dll.insert(500, 5)
	dll.displayDLL()

	dll.delete(500)
	dll.displayDLL()

	for i := 0; i < 5; i++ {
		dll.deleteHead()
		dll.displayDLL()
	}

	for i := 0; i < 5; i++ {
		dll.deleteTail()
		dll.displayDLL()
	}
}
