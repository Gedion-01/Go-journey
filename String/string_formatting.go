package main

import "fmt"

func main() {
	// fmt.Printf :
	// This function is used for formatting strings and printing them to standard output (usually the console)
	// %v : we can use this if we are not assured what else to use
	fmt.Printf("I am %v years old\n", 10)

	// %s : interpolate a string
	fmt.Printf("I am %s years old\n", "way to many")

	// %d : interpolate an integer
	fmt.Printf("I am %d years old\n", 10)

	// %f : interpolate a decimal
	fmt.Printf("I am %f years old\n", 10.5)

	// the ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old\n", 10.5)

	// fmt.Sprintf :
	// This function is similar to fmt.Printf, but instead of printing the formatted string to standard output,
	// it returns the formatted string. It accepts a format string followed by any number of arguments
	str := fmt.Sprintf("Hello %s", "World!")
	fmt.Println(str)
}
