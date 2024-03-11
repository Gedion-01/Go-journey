package main

import "fmt"

func main() {
	accountage := 2.6

	// cast
	accountAgeInt := int(accountage)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
