// some dumb flag
package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	fmt.Printf("%-16s: %b\n", "FlagUp", FlagUp)
	fmt.Printf("%-16s: %b\n", "FlagBroadcast", FlagBroadcast)
	fmt.Printf("%-16s: %b\n", "FlagLoopback", FlagLoopback)
	fmt.Printf("%-16s: %b\n", "FlagPointToPoint", FlagPointToPoint)
	fmt.Printf("%-16s: %b\n", "FlagMulticast", FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)
