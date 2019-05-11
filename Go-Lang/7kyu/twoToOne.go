package main

import "strings"

/*
Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters,

each taken only once - coming from s1 or s2.
Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/
func twoToOne(s1 string, s2 string) string {
	chars := map[byte]int{
		'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0,
		'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0,
		'w': 0, 'x': 0, 'y': 0, 'z': 0,
	}

	tempStr := strings.TrimSpace(s1 + s2)
	if len(tempStr) == 0 {
		return ""
	}

	for i := 0; i < len(tempStr); i++ {
		chars[tempStr[i]]++
	}

	result := []byte{}
	for i := byte('a'); i <= byte('z'); i++ {
		if chars[i] > 0 {
			result = append(result, i)
		}
	}
	return string(result)
}
