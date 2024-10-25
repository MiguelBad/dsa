package main

import "fmt"

func binarySearch(arr []int, needle int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		middle := (high-low)/2 + low

		if arr[middle] == needle {
			return middle
		} else if arr[middle] > needle {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}

func printResult(idx int, needle int) {
	if idx == -1 {
		fmt.Printf("%d is not found\n", needle)
	} else {
		fmt.Printf("%d is found in index %d\n", needle, idx)
	}
}

func main() {
	arr := []int{3, 6, 8, 10, 32, 38}

	needle := 10
	result := binarySearch(arr, needle)
	printResult(result, needle)

	needle = 9
	result = binarySearch(arr, needle)
	printResult(result, needle)
}
