package main

import (
	"./pkgtwo"
	"fmt"
	"os"
	"strings"
)

func fun1() {
	fmt.Println("fun1 print")
}

func main() {
	fmt.Println(len(os.Args))
	for _,arg :=range os.Args{
		fmt.Println(arg)
	}
	fmt.Println("hello go world")

	fmt.Println(pkgtwo.A)

	fmt.Println(os.Args)


	pkgtwo.Print2("b")

	fmt.Println(strings.Join(os.Args," "))



	fun1()
}
