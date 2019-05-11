package main

import "math"

/*
Grasshopper - Terminal game combat function
------------------------------------------------
Create a combat function that takes the player's current health and the amount of damage recieved,
and returns the player's new health. Health can't be less than 0.
*/
func combat(health, damage float64) float64 {
	if math.Floor(damage) == 100.0 {
		return 0
	}

	health = health * damage / 100.0
	if health < 1.1 {
		return 0
	}

	return health
}
