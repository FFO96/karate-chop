package main

import "fmt"

func main() {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10, 12, 13, 14, 15, 16, 17, 60, 200, 400, 500}
	target := 4
	first := 0
	last := len(orderedArray) - 1

	fmt.Println(iterChop(target, orderedArray))
	fmt.Println(recursiveChop(target, first, last, orderedArray))
	fmt.Println(pRecursiveChop(&target, &first, &last, &orderedArray))

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

// Third implementation: Recursive method using pointers
func pRecursiveChop(target, first, last *int, orderedArray *[]int) int {

	if *first <= *last {

		mid := (*first + *last) / 2

		if (*orderedArray)[mid] > *target {
			*last = mid - 1
			return pRecursiveChop(target, first, last, orderedArray)
		} else if (*orderedArray)[mid] < *target {
			*first = mid + 1
			return pRecursiveChop(target, first, last, orderedArray)
		} else {
			return mid
		}
	}
	return -1
}
