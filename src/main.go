package main

import "fmt"

func main() {

	fmt.Println(iterChop(50, []int{1, 3, 4, 5, 8, 9, 10}))

}

// First implementation: Iteration method
func iterChop(target int, orderedArray []int) int {

	first := 0
	last := len(orderedArray) - 1

	for first <= last {

		mid := (first + last) / 2

		if orderedArray[mid] > target {
			last = mid - 1
		} else if orderedArray[mid] < target {
			first = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

/*
func recursiveChop(target int, orderedArray []int) int {
	return -1
}*/
