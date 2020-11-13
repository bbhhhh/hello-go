package main

import "fmt"

const (
	//z = iota
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func main() {
	fmt.Println(KiB, MiB, GiB)
}
