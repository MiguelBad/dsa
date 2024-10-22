package main

import (
	"fmt"
)

type MinHeap struct {
	arr    []int
	length int
}

func (h *MinHeap) insert(item int) {
	h.arr = append(h.arr, item)
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) extractMin() *int {
	if h.length == 0 {
		fmt.Println("Heap is empty")
		return nil
	}

	item, arr := h.arr[0], h.arr[1:]
	if arr[0] > arr[1] {
		arr[1], arr[0] = arr[0], arr[1]
	}

	h.arr = arr
	h.length--
	return &item
}

func (h *MinHeap) delete(item int) *int {
	if h.length == 0 {
		fmt.Println("Heap is empty")
		return nil
	}

	currIdx := 0
	for currIdx < h.length {
		leftIdx, righIdx := _getLeftIdx(currIdx), _getRightIdx(currIdx)

		if h.arr[currIdx] == item {
			item, arr := h.arr[currIdx], h.arr[currIdx:]
		}

		if h.arr[currIdx] > item {
			currIdx = leftIdx
		} else {
			currIdx = righIdx
		}
	}

	h.length--
	return
}

func deleteHelper() {
}

func (h *MinHeap) heapifyUp(currIdx int) {
	parentIdx := _getParentIdx(currIdx)

	if h.arr[parentIdx] > h.arr[currIdx] {
		h.arr[parentIdx], h.arr[currIdx] = h.arr[currIdx], h.arr[parentIdx]
		h.heapifyUp(parentIdx)
	} else {
		return
	}
}

func _getParentIdx(currIdx int) int {
	return (currIdx - 1) / 2
}

func _getLeftIdx(currIdx int) int {
	return currIdx*2 + 1
}

func _getRightIdx(currIdx int) int {
	return currIdx*2 + 2
}

func main() {
	heap := &MinHeap{
		arr:    []int{},
		length: 0,
	}

	heap.insert(3)
	heap.insert(9)
	heap.insert(7)
	heap.insert(1)
	heap.insert(3)
	heap.insert(0)
	heap.insert(3)
	heap.insert(3)
	heap.insert(3)

	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()

	fmt.Println(heap.arr)
}
