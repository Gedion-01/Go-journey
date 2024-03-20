package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}

	process(data, printValue)
}

func process(data []int, callback func(int)) {
	for _, value := range data {
		callback(value)
	}
}
func printValue(value int) {
	fmt.Println("Value:", value)
}
