// Package tmep conv performs Celsius and Fahrenheit temparture computations.
package tempconv

import "fmt"

// Celsius is a good thing.
type Celsius float64

// Fahrenheit is a good thing.
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -237.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
