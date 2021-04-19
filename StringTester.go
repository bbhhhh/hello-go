package main

import "fmt"

func main() {
	var a string = "abc"
	a = "def   中话   def  gsdg"
	fmt.Println(a)

	for _, r := range a {
		fmt.Printf("\r%c", r)
	}
	for _, r := range "-\\|/" {
		fmt.Printf("\r%c", r)
	}

}
