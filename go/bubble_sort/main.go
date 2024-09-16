package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{10, 3, 4, 5, 8, 1, 9}
	fmt.Println("Unsorted: ", arr)

	bubbleSort(arr)

	fmt.Println("Sorted: ", arr)
}
