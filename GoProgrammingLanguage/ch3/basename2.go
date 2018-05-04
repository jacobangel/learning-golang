// base name stuff
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename2("abc"))
	fmt.Println(comma("asdfasdfasdfasdfafsd"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // negative one is good here i guess
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
