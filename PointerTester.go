package main

import "fmt"

type Kerwen float64
type Celsius float64

const (
	AbsZeorC Celsius = -273.15
	ZeorK            = Kerwen(AbsZeorC)
)

func CtoK(c Celsius) Kerwen {
	return Kerwen(c) - ZeorK
}

func (k Kerwen) String() string {
	return fmt.Sprintf("%gK", k)
}

func main() {
	j := t1(10)

	fmt.Println(*j)

	fmt.Println(gcd(11, 20))

	type t1 int32
	type t2 int64

	var i1 t1 = 1
	var i2 int = 3
	var i3 float32 = 1.0
	i2 = int(i1)
	i2 = int(i3)
	fmt.Println(i1, i2, i3)
	fmt.Println(CtoK(Celsius(i2)).String())
	p1 := &i1
	*p1 = 20
	fmt.Println(*p1)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func t1(i int) *int {
	return &i
}
