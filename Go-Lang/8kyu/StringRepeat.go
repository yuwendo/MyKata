package main

import "strings"

/*
String repeat
-------------------------------
Write a function called repeatStr which repeats the given string string exactly n times.

	repeatStr(6, "I") // "IIIIII"
	repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"
	repeatStr(6, "I") -- "IIIIII"
	repeatStr(5, "Hello") -- "HelloHelloHelloHelloHello"
	repeat_str(6, "I") # "IIIIII"
	repeat_str(5, "Hello") # "HelloHelloHelloHelloHello"
	StringRepeat.repeatStr(6, "I") # "IIIIII"
	StringRepeat.repeatStr(5, "Hello") # "HelloHelloHelloHelloHello"
	repeatStr(6, "I") # "IIIIII"
	repeatStr(5, "Hello") # "HelloHelloHelloHelloHello"
	repeatstr(6, "I") # "IIIIII"
	repeatstr(5, "Hello") # "HelloHelloHelloHelloHello"
	repeatstr(6, "I") # "IIIIII"
	repeatstr(5, "Hello") # "HelloHelloHelloHelloHello"
*/
func repeatStr(repititions int, value string) string {
	var result []string
	for times := 0; times < repititions; times++ {
		result = append(result, value)
	}

	return strings.Join(result, "")
}
