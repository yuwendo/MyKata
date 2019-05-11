package main

import "sort"

/*
Two Oldest Ages
-------------------------------
The two oldest ages function/method needs to be completed.
It should take an array of numbers as its argument and return the two highest numbers within the array.
The returned value should be an array in the format [second oldest age, oldest age].

The order of the numbers passed in could be any order.
The array will always include at least 2 items.

For example:
	twoOldestAges( [1, 2, 10, 8] ) // should return [8, 10]
	twoOldestAges [34, 45, 12, 98, 0, 23] # should return [45, 98]
	two_oldest_ages([1, 3, 10, 0]) # should return [3, 10]
	TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}
	twoOldestAges({1, 5, 87, 45, 8, 8}) -- should return {45, 87}
	two_oldest_ages([1, 5, 87, 45, 8, 8]) # should return [45, 87]
	LargestTwo.TwoOldestAges(new int[] {1, 2, 10, 8}) => new int[] {8, 10}
	twoOldestAges(listOf(1, 5, 87, 45, 8, 8)) // should return listOf(45, 87)
	twoOldestAges(@[1, 5, 87, 45, 8, 8]) # should return [45, 87]
	two_oldest_ages([1, 3, 10, 0]) # should return [3, 10]
*/
func twoOldestAges(ages []int) [2]int {
	oldAges := make([]int, 2)

	for i := 0; i < len(ages); i++ {
		if ages[i] > oldAges[0] {
			oldAges[0] = ages[i]
			sort.Ints(oldAges)
		}
	}

	result := [2]int{oldAges[0], oldAges[1]}

	return result
}
