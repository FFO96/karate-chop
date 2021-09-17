# Karate-chop
A binary chop (sometimes called the more prosaic binary search) finds the position of value in a sorted array of values. It achieves some efficiency by halving the number of items under consideration each time it probes the values: in the first pass it determines whether the required value is in the top or the bottom half of the list of values. In the second pass in considers only this half, again dividing it in to two. It stops when it finds the value it is looking for, or when it runs out of array to search

This program needs to be set with an int [target](#target) and an [SIA](#sia) . It will return the integer index of the target in the SIA, or -1 if the target is not in the SIA.

## Problem solving
To solve the problem we can differentiate between two different implementations: 

#### Iteration method
The traditional iterative approach involves using a loop to modify the inferior and superior indexes until the target is found in the SIA.

#### Recursive method
The recursive approach also involves modifying the inferior and superior indexes but instead of using a loop a recursive function is used.

When solving the problem, we have distinguished two implementations but in the code we find three. This is because the third method is an evolution of the recursive implementation. The parameters of this third function are pointers to the required variables in order to save memory during the recursion.

## Keywords

- <a id="target">**Target**</a>: Int value to search in the SIA.
- <a id="sia">**SIA**</a>: Sorted int array.

## Run the progam

- **Normal program**</a>: go run src/main.go
- **Testing**</a>: go test src/main_test.go src/main.go -v

## Reference
Word-chain problem: [http://codekata.com/kata/kata02-karate-chop/](http://codekata.com/kata/kata02-karate-chop/)
