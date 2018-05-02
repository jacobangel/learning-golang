package main

import (
	"fmt"
	"os"
)

func main() {
	sAcc, sep := "", ""
	for ind, arg := range os.Args[1:] {
		// interesting here, because the arg gets nulled
		// and seen as a falsy value evenrtually.
		sAcc += fmt.Sprintf("%d %s%s", ind, arg, sep)
		sep = ", "
	}
	fmt.Println(sAcc)
}
