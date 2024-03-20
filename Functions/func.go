package main

import "fmt"

func main() {
	fmt.Println(sub(10, 5))
	fmt.Println(add(10, 2))
}

func sub(x int, y int) int {
	return x - y
}
func add(x, y int) int {
	return x + y
}
