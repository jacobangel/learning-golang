// Temp user stuff
package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	boilC := tempconv.FToC(boilingF)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC)
	fmt.Println(boilC.String())
	fmt.Println(boilingF.String())
	fmt.Println(tempconv.FToK(boilingF))
	fmt.Println(tempconv.CToK(boilC))
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
}
