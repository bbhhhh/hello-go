package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "abcsdf士大夫as作为法定asd"
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
}
