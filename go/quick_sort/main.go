package main

import "fmt"

func partition(arr []int, low int, high int) int {
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] < arr[high] {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[high], arr[idx] = arr[idx], arr[high]
	return idx
}

func quickSort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	pivot_idx := partition(arr, low, high)

	quickSort(arr, low, pivot_idx-1)
	quickSort(arr, pivot_idx+1, high)
}

func main() {
	arr := []int{29, 14, 1, 49, 32, 50, 2, 14}
	low := 0
	high := len(arr) - 1

	fmt.Println("Unsorted: ", arr)
	quickSort(arr, low, high)
	fmt.Println("Sorted: ", arr)
}
