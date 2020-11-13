package main

import "fmt"

func main() {
	aa := "adasdf"
	bb := "中国"
	fmt.Println(len(aa), len(bb))
	for i := 0; i < len(aa); i++ {
		fmt.Println(len(aa[:i]))
	}
}
