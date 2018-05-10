package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow and extend the slice
		z = x[:zlen]
	} else {
		// no dize, need to double it.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	fmt.Println(appendInt(a, 7))
	fmt.Println(appendInt(b, 8))

	var x, y []int
	fmt.Println("AppendInt")
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var z []int
	fmt.Println("Append")
	for i := 0; i < 10; i++ {
		y = append(z, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		z = y
	}
}
