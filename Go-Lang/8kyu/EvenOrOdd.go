package main

/*
Even or Odd
----------------------------
Create a function (or write a script in Shell) that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/
func evenOrOdd(number int) string {
	// your code here
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
