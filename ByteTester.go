package main

import "fmt"

func main() {
	//var i int = 256
	var b byte = 255
	fmt.Println(b)
	aa := "abc"
	var slice []byte
	slice = []byte(aa)
	slice[9] = 1
	fmt.Println(slice)
}
