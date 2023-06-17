package main

import (
	"fmt"
	"math"
)

/*
GenDisplaceFn generates a function that computes displacement based on the given acceleration,
initial velocity, and initial displacement
*/
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	// Prompt the user to enter values for acceleration, initial velocity, and initial displacement
	fmt.Println("Enter acceleration (m/s^2):")
	fmt.Scanln(&acceleration)
	fmt.Println("Enter initial velocity (m/s):")
	fmt.Scanln(&initialVelocity)
	fmt.Println("Enter initial displacement (m):")
	fmt.Scanln(&initialDisplacement)

	// Generate the displacement function
	displacementFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Prompt the user to enter a value for time
	fmt.Println("Enter time (s):")
	fmt.Scanln(&time)

	// Calculate and print the displacement at the entered time
	displacement := displacementFn(time)
	fmt.Println("Calculated displacement (m):", displacement)
}
