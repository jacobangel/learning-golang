// echo stuff
package main

import (
	"flag"
	"fmt"
	"strings"
)

// oh this is a super super neat package.
var n = flag.Bool("n", false, "omit trailing new line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
