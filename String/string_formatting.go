package main

import "fmt"

func main() {
	// fmt.Printf
	// %v : we can use this if we are not assured what else to use
	fmt.Printf("I am %v years old\n", 10)

	// %s : interpolate a string
	fmt.Printf("I am %s years old\n", "way to many")

	// %d : interpolate an integer
	fmt.Printf("I am %d years old\n", 10)

	// %f : interpolate a decimal
	fmt.Printf("I am %f years old\n")
	// fmt.Sprintf

}
