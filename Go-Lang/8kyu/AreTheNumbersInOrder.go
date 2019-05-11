package main

import "sort"

/*
Are the numbers in order?
---------------------------
In this Kata, your function receives an array of integers as input.
Your task is to determine whether the numbers are in ascending order.
An array is said to be in ascending order if there are no two adjacent integers where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. non-empty arrays containing only integers.

Note that an array of 1 integer is automatically considered to be sorted in ascending order
since all (zero) adjacent pairs of integers satisfy the condition that the left integer does not exceed the right integer in value.

An empty list is considered a degenerate case and therefore will not be tested in this Kata -
	feel free to raise an Issue if you see such a list being tested.

For example:
	isAscOrder(new int[]{1,2,4,7,19}) == true
	isAscOrder(new int[]{1,2,3,4,5}) == true
	isAscOrder(new int[]{1,6,10,18,2,4,20}) == false
	isAscOrder(new int[]{9,8,7,6,5,4,3,2,1}) == false // numbers are in DESCENDING order
	in_asc_order({1,2,4,7,19}, 5); // returns true
	in_asc_order({1,2,3,4,5}, 5); // returns true
	in_asc_order({1,6,10,18,2,4,20}, 7); // returns false
	in_asc_order({9,8,7,6,5,4,3,2,1}, 9); // returns false because the numbers are in DESCENDING order
	inAscOrder([1,2,4,7,19]); // returns true
	inAscOrder([1,2,3,4,5]); // returns true
	inAscOrder([1,6,10,18,2,4,20]); // returns false
	inAscOrder([9,8,7,6,5,4,3,2,1]); // returns false because the numbers are in DESCENDING order
	inAscOrder([1,2,4,7,19]); // returns true
	inAscOrder([1,2,3,4,5]); // returns true
	inAscOrder([1,6,10,18,2,4,20]); // returns false
	inAscOrder([9,8,7,6,5,4,3,2,1]); // returns false because the numbers are in DESCENDING order
	Kata.IsAscOrder(new int[]{1,2,4,7,19}) == true
	Kata.IsAscOrder(new int[]{1,2,3,4,5}) == true
	Kata.IsAscOrder(new int[]{1,6,10,18,2,4,20}) == false
	Kata.IsAscOrder(new int[]{9,8,7,6,5,4,3,2,1}) == false // numbers are in DESCENDING order
	in_asc_order([1,2,4,7,19]) # returns True
	in_asc_order([1,2,3,4,5]); // returns True
	in_asc_order([1,6,10,18,2,4,20]) # returns False
	in_asc_order([9,8,7,6,5,4,3,2,1]) # returns False because the numbers are in DESCENDING order
	in_asc_order([1, 2, 4, 7, 19]); // true
	in_asc_order([1, 2, 3, 4, 5]); // true
	in_asc_order([1, 6, 10, 18, 2, 4, 20]); // false
	in_asc_order([9, 8, 7, 6, 5, 4, 3, 2, 1]); // false (NOTE: because the numbers are in DESCENDING order, not ascending order)
	isAscOrder [1, 2, 4, 7, 19] -- True
	isAscOrder [1, 2, 3, 4, 5] -- True
	isAscOrder [1, 6, 10, 18, 2, 4, 20] -- False
	isAscOrder [9, 8, 7, 6, 5, 4, 3, 2, 1] -- False (NOTE: because the numbers are in DESCENDING order, not ascending order)
	is_asc_order([1,2,4,7,19]) # returns True
	is_asc_order([1,2,3,4,5]); // returns True
	is_asc_order([1,6,10,18,2,4,20]) # returns False
	is_asc_order([9,8,7,6,5,4,3,2,1]) # returns False because the numbers are in DESCENDING order
	is_asc_order([1,2,4,7,19]) # returns True
	is_asc_order([1,2,3,4,5]); // returns True
	is_asc_order([1,6,10,18,2,4,20]) # returns False
	is_asc_order([9,8,7,6,5,4,3,2,1]) # returns False because the numbers are in DESCENDING order
	InAscOrder([]int{1, 2, 4, 7, 19}) # returns True
	InAscOrder([]int{1, 2, 3, 4, 5}) // returns True
	InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}) # returns False
	InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) # returns False because the numbers are in DESCENDING order
	in_asc_order({1,2,4,7,19}, 5); // => 1
	in_asc_order({1,2,3,4,5}, 5); // => 1
	in_asc_order({1,6,10,18,2,4,20}, 7); // => 0
	in_asc_order({9,8,7,6,5,4,3,2,1}, 9); // => 0
*/
func inAscOrder(numbers []int) bool {
	return sort.IntsAreSorted(numbers)
}
