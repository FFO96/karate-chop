package main

import "fmt"

func main() {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10}
	target := 4
	fmt.Println(iterChop(target, orderedArray))
	fmt.Println(recursiveChop(target, 0, len(orderedArray)-1, orderedArray))

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

// Second implementation: Recursive method
func recursiveChop(target, first, last int, orderedArray []int) int {

	if first <= last {

		mid := (first + last) / 2

		if orderedArray[mid] > target {
			return recursiveChop(target, first, mid-1, orderedArray)
		} else if orderedArray[mid] < target {
			first = mid + 1
			return recursiveChop(target, mid+1, last, orderedArray)
		} else {
			return mid
		}
	}
	return -1
}
