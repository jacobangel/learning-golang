// Demonstrates pointer stuff
package main

import "fmt"

func main() {
	x := 1  // x is int
	p := &x // p is *int, & derefences
	fmt.Println(*p)
	*p = 3
	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&x)
	incr(&x)
	incr(p)
	fmt.Println(x)
}

func incr(p *int) int {
	*p++
	return *p
}
