package main

/*
Convert boolean values to strings 'Yes' or 'No'.
-------------------------------------------------
Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.
*/
func boolToWord(word bool) string {
	if word == true {
		return "Yes"
	}
	return "No"
}
