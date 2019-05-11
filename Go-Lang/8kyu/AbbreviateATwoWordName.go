package main

import "strings"

/*
Abbreviate a Two Word Name
-----------------------------------------
Write a function to convert a name into initials.
This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot seperating them.

It should look like this:

Sam Harris => S.H

Patrick Feeney => P.F
*/
func abbrevName(name string) string {
	//your code here
	splitStrs := strings.Split(name, " ")
	if len(splitStrs) != 2 {
		return ""
	}
	return strings.ToUpper(splitStrs[0][0:1] + "." + splitStrs[1][0:1])
}
