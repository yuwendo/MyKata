package main

import (
	"math"
	"sort"
)

/*
Check if a triangle is an equable triangle!
------------------------------------------------
A triangle is called an equable triangle if its area equals its perimeter.
Return true, if it is an equable triangle, else return false.
You will be provided with the length of sides of the triangle. Happy Coding!
*/
func equableTriangle(a, b, c int) bool {
	// Your code goes here
	params := []int{a, b, c}
	sort.Ints(params)

	if (params[0] + params[1]) > params[2] {
		s := (a + b + c) / 2
		area := float64(s * (s - a) * (s - b) * (s - c))
		area = math.Sqrt(area)
		if area == float64(a+b+c) {
			return true
		}
	}

	return false
}
