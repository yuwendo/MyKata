package main

/*
Beginner - Lost Without a Map
-------------------------------
Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]

For the beginner, try to use the map method - it comes in very handy quite a lot so is a good one to know.
*/
func maps(x []int) []int {
	x1 := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		x1[i] = x[i] * 2
	}
	return x1
}
