package main

import "fmt"

func fibonacci(num int) int {
	seen := make(map[int]int)
	return fibHelper(num, seen)
}

func fibHelper(num int, seen map[int]int) int {
	if val, ok := seen[num]; ok {
		return val
	}

	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	result := fibHelper(num-1, seen) + fibHelper(num-2, seen)
	seen[num] = result
	return result
}

func displayText(num int, answer int) {
	fmt.Printf("The %vth fibonacci is %v\n", num, answer)
}

func main() {
	for _, i := range []int{3, 5, 7, 9} {
		answer := fibonacci(i)
		displayText(i, answer)
	}
}

