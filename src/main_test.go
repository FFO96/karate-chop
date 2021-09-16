package main

import (
	"testing"
)

func TestIterChop(t *testing.T) {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10, 12, 13, 14, 15, 16, 17, 60, 200, 400, 500}
	target := 4

	// Setting the expected index for the target "4"
	expectedIndex := 2

	// Comparing solution and expected solution
	if expectedIndex != iterChop(target, orderedArray) {
		t.Fail()
	}

	// Setting the expected index for the target "3"
	target = 3
	expectedIndex = -1

	// Comparing solution and expected solution
	if expectedIndex != iterChop(target, orderedArray) {
		t.Fail()
	}

}

func TestRecursiveChop(t *testing.T) {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10, 12, 13, 14, 15, 16, 17, 60, 200, 400, 500}
	target := 4
	first := 0
	last := len(orderedArray) - 1

	// Setting the expected index for the target "4"
	expectedIndex := 2

	// Comparing solution and expected solution
	if expectedIndex != recursiveChop(target, first, last, orderedArray) {
		t.Fail()
	}

	// Setting the expected index for the target "3"
	target = 3
	expectedIndex = -1

	// Comparing solution and expected solution
	if expectedIndex != recursiveChop(target, first, last, orderedArray) {
		t.Fail()
	}

}

func TestPRecursiveChop(t *testing.T) {

	orderedArray := []int{1, 2, 4, 5, 8, 9, 10, 12, 13, 14, 15, 16, 17, 60, 200, 400, 500}
	target := 4
	first := 0
	last := len(orderedArray) - 1

	// Setting the expected index for the target "4"
	expectedIndex := 2

	// Comparing solution and expected solution
	if expectedIndex != pRecursiveChop(&target, &first, &last, &orderedArray) {
		t.Fail()
	}

	// Setting the expected index for the target "3"
	target = 3
	expectedIndex = -1

	// Comparing solution and expected solution
	if expectedIndex != pRecursiveChop(&target, &first, &last, &orderedArray) {
		t.Fail()
	}

}
