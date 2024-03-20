package main

import "fmt"

func getPoint() (x int, y int) {
	return 3, 4
}

func main() {
	x, _ := getPoint() // we are ignoring y
	fmt.Println(x)
}
