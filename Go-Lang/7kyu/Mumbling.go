package main

import "strings"

/*
Mumbling
-------------------------
This time no story, no theory.
The examples below show you how to write function accum:

Examples:
	accum("abcd") -> "A-Bb-Ccc-Dddd"
	accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
	accum("cwAt") -> "C-Ww-Aaa-Tttt"
	The parameter of accum is a string which includes only letters from a..z and A..Z.
*/
func accum(s string) string {
	// your code
	var result []string
	s = strings.ToLower(s)
	for str := 0; str < len(s); str++ {
		tempStr := strings.Repeat(s[str:str+1], str+1)
		result = append(result, strings.Title(tempStr))
	}
	return strings.Join(result, "-")
}
