package main

/*
Is this a triangle?
------------------------------
Implement a method that accepts 3 integer values a, b, c.
The method should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).
*/
func isTriangle(a, b, c int) bool {
	if (a <= 0) || (b <= 0) || (c <= 0) {
		return false
	}

	if ((a + b) > c) && ((b + c) > a) && ((a + c) > b) {
		return true
	}

	return false
}
