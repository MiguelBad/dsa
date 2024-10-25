package main

import "fmt"

// heap sort for min heap
func heapSort(arr []int) {
	for i := len(arr)/2 - 2; i > -1; i-- {
		heapSortHelper(arr, i)
	}
}

func heapSortHelper(arr []int, idx int) {
	left := leftIdx(idx)
	right := rightIdx(idx)
	smallest := idx

	if left < len(arr) && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest == idx {
		return
	}

	arr[smallest], arr[idx] = arr[idx], arr[smallest]
	heapSortHelper(arr, smallest)
}

func leftIdx(idx int) int {
	return idx*2 + 1
}

func rightIdx(idx int) int {
	return idx*2 + 2
}

func main() {
	arr := []int{49, 9, 2, 16, 6, 3}
	fmt.Printf("Original arr:\n\t%v\n", arr)
	heapSort(arr)
	fmt.Printf("Sorted arr:\n\t%v\n", arr)
}
