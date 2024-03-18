package main

import "fmt"

func main() {
	var p =  f()
	fmt.Println(*p)
	// each call of f returns distinct value
	fmt.Println(f() == f()) // false
	v := 1

	incr(&v) // side effect
	incr(&v)
	fmt.Println(v)
}

func f() *int {
	v := 12
	return &v
}
func incr( p *int) int {
	*p++
	return *p
}
