package main

import "fmt"

func main() {
	s1 := []int{1}
	fmt.Printf("%T %[1]v\n", s1)
	s2 := make([]int, 1, 10)
	fmt.Printf("%q %v %v\n", s2, len(s2), cap(s2))
}
