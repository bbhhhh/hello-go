package main

import "fmt"

func main() {
	a1 := [...]int{1, 2}
	for _, v := range a1 {
		fmt.Println(v)
	}
	a1 = [2]int{3}
	a2 := [...]int{0: 0, 2: 2, 1: 1, 7: 7}
	fmt.Println(a2)
	fmt.Printf("%T %[1]v\n", a2)
	a2[5] = 4
	a3 := a2[3:4]
	fmt.Println(len(a3), cap(a3))
	a3[0] = 10
	a4 := make([]int, 10, 20)
	fmt.Printf("%T %[1]v\n", a4)
	a5 := []int{1, 2}
	fmt.Println(a5)

}
