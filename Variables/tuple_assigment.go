package main

import "fmt"

func main() {
	//
	fmt.Println(fib(6))
}

func fib(n int) int {
	x, y := 0, 1
	fmt.Println("initial x & y: ", x, y)
	for i := 0; i < n; i++ {
		fmt.Println("=> ", x, "+", y)
		x, y = y, x+y
		fmt.Println("x & y result: ", x, y)
	}
	return x
}
