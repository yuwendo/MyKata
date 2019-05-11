package main

/*
Reversed Strings (Go)
---------------------------
Complete the solution so that it reverses the string value passed into it.

Example:
	solution('world'); // returns 'dlrow'
*/
func solution(word string) string {
	if len(word) < 2 {
		return word
	}

	result := ""
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}

	return result
}
