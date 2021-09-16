package main

import "fmt"

func main() {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10, 12, 13, 14, 15, 16, 17, 60, 200, 400, 500}
	target := 4
	first := 0
	last := len(orderedArray) - 1

	fmt.Printf("Iterative chop result: %v\n", iterChop(target, orderedArray))
	fmt.Printf("Recursive chop result: %v\n", recursiveChop(target, first, last, orderedArray))
	fmt.Printf("Recursive chop with pointers result: %v\n", pRecursiveChop(&target, &first, &last, &orderedArray))
}

// First implementation: Iteration method
func iterChop(target int, orderedArray []int) int {

	first := 0
	last := len(orderedArray) - 1

	// While the inferior index is not bigger than the superior index, it will keep searching
	for first <= last {

		// Calculating array middle point
		mid := (first + last) / 2

		// Target less than number in middle point
		if orderedArray[mid] > target {
			last = mid - 1

			// Target greater than number in middle point
		} else if orderedArray[mid] < target {
			first = mid + 1
		} else {
			// Target found
			return mid
		}
	}
	// Target not in array
	return -1
}

// Second implementation: Recursive method
func recursiveChop(target, first, last int, orderedArray []int) int {

	if first <= last {

		// While the inferior index is not bigger than the superior index, it will keep searching
		mid := (first + last) / 2

		// Target less than number in middle point
		if orderedArray[mid] > target {
			return recursiveChop(target, first, mid-1, orderedArray)

			// Target greater than number in middle point
		} else if orderedArray[mid] < target {
			first = mid + 1
			return recursiveChop(target, mid+1, last, orderedArray)
		} else {
			// Target found
			return mid
		}
	}
	// Target not in array
	return -1
}

// Third implementation: Recursive method using pointers
func pRecursiveChop(target, first, last *int, orderedArray *[]int) int {

	// While the inferior index is not bigger than the superior index, it will keep searching
	if *first <= *last {

		mid := (*first + *last) / 2

		// Target less than number in middle point
		if (*orderedArray)[mid] > *target {
			*last = mid - 1
			return pRecursiveChop(target, first, last, orderedArray)

			// Target greater than number in middle point
		} else if (*orderedArray)[mid] < *target {
			*first = mid + 1
			return pRecursiveChop(target, first, last, orderedArray)
		} else {
			// Target found
			return mid
		}
	}
	// Target not in array
	return -1
}
