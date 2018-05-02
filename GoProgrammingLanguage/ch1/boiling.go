// Boiling piont printer
package main

import "fmt"

// package level declaration
const boilingF = 212.0

func main() {
	// local declarations
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g° or %gC°\n", f, c)
	// Ouput
	// boiling point = 212 or 100c
}
