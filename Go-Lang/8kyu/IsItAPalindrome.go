package main

import "strings"

/*
Write function isPalindrome that checks if a given string (case insensitive) is a palindrome.
*/
func isPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	}

	str = strings.ToLower(str)
	if str[0] != str[len(str)-1] {
		return false
	}

	for i := 1; i < len(str)/2; i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}

	return true
}
