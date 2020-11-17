package main

import (
	"fmt"
	"io"
	"os"
)

type ShapeI interface {
	ShapeName() string
}

type ShapeII interface {
	ShapeName() string
}

type Shape string

func (s Shape) ShapeName() string {
	return string(s)
}

func main() {
	s1 := Shape("triangle")
	var s2 ShapeI
	var s3 ShapeII
	s2 = s1
	s3 = s1
	s2 = s3
	fmt.Println(s2.ShapeName())
	fmt.Println(s3.ShapeName())

	var w io.Writer
	w = os.Stderr
	w.Write([]byte("hello byte"))
}
