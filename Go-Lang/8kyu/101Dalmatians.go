package main

/*
101 Dalmatians - squash the bugs, not the dogs!
------------------------------------------------
Your friend has been out shopping for puppies (what a time to be alive!)...
He arrives back with multiple dogs, and you simply do not know how to respond!

By repairing the function provided, you will find out exactly how you should respond, depending on the number of dogs he has.

The number of dogs will always be a number and there will always be at least 1 dog.
*/
func howManyDalmatians(number int) string {
	if number < 11 {
		return "Hardly any"
	} else if number < 51 {
		return "More than a handful!"
	} else if number < 101 {
		return "Woah that's a lot of dogs!"
	}

	return "101 DALMATIONS!!!"
}
