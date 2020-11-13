package main

import "fmt"

func main() {
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q", newline)
}
