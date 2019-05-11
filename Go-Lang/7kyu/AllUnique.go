package main

/*
All unique
------------------------------
Write a program to determine if a string contains all unique characters.
Return true if it does and false otherwise.

The string may contain any of the 128 ASCII characters.
*/
func hasUniqueChar(str string) bool {
	if len(str) == 0 {
		panic("Null string")
		return false
	}

	tempRune := []rune(str)
	for id, _ := range tempRune {
		for i := 0; i < id; i++ {
			if tempRune[i] == tempRune[id] {
				return false
			}
		}
	}

	return true
}
