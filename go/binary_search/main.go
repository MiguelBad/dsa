package main

import "fmt"

func binarySearch(arr []int, needle int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == needle {
			return mid
		} else if arr[mid] > needle {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 5, 12, 16, 20, 24, 38, 40, 48}

	for _, needle := range []int{5, 1, 33, 48} {
		idx := binarySearch(arr, needle)

		if idx == -1 {
			fmt.Printf("Number: %d\nNot found\n\n", needle)
		} else {
			fmt.Printf("Number: %d\nAt index: %d\n\n", needle, idx)
		}
	}
}
