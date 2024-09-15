package main

import "fmt"

func reverse(arr []int) {
	last := len(arr) - 1

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[last-i] = arr[last-i], arr[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	reverse(arr)

	fmt.Println(arr)
}
