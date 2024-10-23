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
	h.length++
	_heapifyUp(h.length-1, h.arr)
}

func _heapifyUp(currIdx int, arr []int) {
	if currIdx < 1 {
		return
	}

	parentIdx := _getParentIdx(currIdx)

	if arr[parentIdx] > arr[currIdx] {
		arr[parentIdx], arr[currIdx] = arr[currIdx], arr[parentIdx]
		_heapifyUp(parentIdx, arr)
	} else {
		return
	}
}

func (h *MinHeap) extractMin() *int {
	if h.length == 0 {
		fmt.Println("Heap is empty")
		return nil
	}

	h.length--
	h.arr[h.length], h.arr[0] = h.arr[0], h.arr[h.length]
	item := h.arr[h.length]
	h.arr = h.arr[:h.length]
	_heapifyDown(0, h.arr)
	return &item
}

func _heapifyDown(currIdx int, arr []int) {
	leftIdx, rightIdx := _getLeftIdx(currIdx), _getRightIdx(currIdx)
	smallest := currIdx

	if leftIdx < len(arr)-1 && arr[leftIdx] < arr[smallest] {
		smallest = leftIdx
	}

	if rightIdx < len(arr)-1 && arr[rightIdx] < arr[smallest] {
		smallest = rightIdx
	}

	if smallest != currIdx {
		arr[smallest], arr[currIdx] = arr[currIdx], arr[smallest]
		_heapifyDown(smallest, arr)
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
	heap.insert(1)
	heap.insert(4)
	heap.insert(2)

	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
	heap.extractMin()
	fmt.Println(heap.arr)
}
