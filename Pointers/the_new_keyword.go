package main

import "fmt"
/ *

*/
func main() {
	// unnamed variable
	p := new(int) // returns its address
	fmt.Println(p)
	*p = 2
	fmt.Println(*p)
}

func newInt() *int {
	
	return new(int)
}
