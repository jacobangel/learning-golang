package main

import (
	"fmt"
)

var months = [...]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	secondQuarter := months[4:7]
	summer := months[6:9]
	fmt.Println(secondQuarter)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range secondQuarter {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	extended := summer[:5]
	fmt.Println(extended)

	a := [...]int{0, 1, 2, 3, 4, 5}
	rev(a[:])
	fmt.Printf("%v\n%T\n", a, a)
}

func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
