// FTOC prints conversions.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	for _, temp := range os.Args[1:] {
		tempF, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("%g°F = %g°C\n", tempF, fToC(tempF))
	}
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
