package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{29, 14, 1, 49, 32, 50, 2, 14}

    fmt.Println("Unsorted", arr)
	bubbleSort(arr)
    fmt.Println("Sorted", arr)
}
