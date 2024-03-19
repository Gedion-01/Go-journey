package main

import "fmt"

func main() {
	height := 5

	if height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}

	// Initial statment of an if block
	// if INITIAL_STATEMENT; CONDITION {}

	email := ""
	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	}
}

func getLength(email string) {
	panic("unimplemented")
}
