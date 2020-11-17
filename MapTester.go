package main

import "fmt"

type AMap map[string]int
type BMap map[int]AMap

var C BMap

func main() {
	ages := make(map[string]int)
	ages["a"] = 0
	ages["b"] = 2
	i, ok := ages["c"]
	fmt.Println(i, ok)
	j := ages["b"]
	fmt.Println(j, ok)

	x := AMap{"a": 1}
	y := BMap{1: x}
	fmt.Println(y)
}
