package main

/*
Sum of positive
-----------------------------------
You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/
func positiveSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < 0 {
			continue
		}
		sum += numbers[i]
	}

	return sum
}
