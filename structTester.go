package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	type s1 struct {
		m1 int
		m2 string
		m3 []string
		m4 map[int]string
	}

	var x s1
	fmt.Println(x)
	x.m3 = make([]string, 1, 3)
	fmt.Println(len(x.m3), cap(x.m3))
	x.m3 = append(x.m3, "", "", "", "", "", "")
	fmt.Println(len(x.m3), cap(x.m3))
	x.m3[1] = "dsd"
	x.m3[3] = "dsdfd"
	x.m3[4] = "dsdasdf"

	x.m4 = make(map[int]string)
	x.m4[0] = "a"

	x.m1 = 20
	var y *s1
	y = &x
	y.m1 = 30
	(*y).m1 = 40
	fmt.Println(x.m1, y.m1, (*y).m3)
	var z *s1
	z = new(s1)

	fmt.Println(z.m1, (*z).m1)
	intlist := []int{1, 10, 3, 5, 2, 4, 6, 9, 8, 7}
	var root *tree
	for _, v := range intlist {
		root = add(root, v)
	}

	walkthroughTree(root)

}

func walkthroughTree(t *tree) {
	if t != nil {
		walkthroughTree(t.left)
		fmt.Printf("value=%v\n", t.value)
		walkthroughTree(t.right)
	}
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
