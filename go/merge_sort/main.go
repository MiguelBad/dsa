package main

import (
	"fmt"
)

func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	for i := 0; i < mid; i++ {
		left[i] = arr[i]
	}

	for j := mid; j < len(arr); j++ {
		right[j-mid] = arr[j]
	}

	mergeSort(left)
	mergeSort(right)

	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

func main() {
	var arr []int = []int{68, 41, 39, 80, 67, 72, 97, 79, 28, 3, 87, 70, 89, 71, 64, 13, 55, 81, 72, 24}

	fmt.Printf("Unsorted: %v\n", arr)
	mergeSort(arr)
	fmt.Printf("Sorted: %v\n", arr)
}
