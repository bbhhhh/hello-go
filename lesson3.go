package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//var a = 0
	//f, e := os.Open("test")
	//fmt.Println(a, f, e)

//	n := flag.Bool("n", false, "newline")
	n := flag.String("n", "ns", "newline")
	sep := flag.String("s", " ", "sep")
	flag.Parse()
	for i:=0;i<len(flag.Args());i++{
		fmt.Println(flag.Args()[i])
	}
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println(*n)
}
