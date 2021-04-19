package main

import "fmt"

func main() {
	s1 := []int{1}
	fmt.Printf("%T %[1]v\n", s1)
	s2 := make([]int, 1, 10)
	s2[0] = 1
	fmt.Printf("%q %v %v\n", s2, len(s2), cap(s2))
	s3 := new([]string)
	fmt.Printf("%q %v %v\n", s3, len(*s3), cap(*s3))
	*s3 = append(*s3, "abc")
	var s4 []string
	s4 = *s3
	s4[0] = "abc"
	fmt.Println(*s3)
}
