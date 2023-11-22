package main

import "fmt"

type Test struct {
	a, b int
}

func (t Test) Abs() {
	t.a = 10
	t.b = 10
}

func (t *Test) PointerAbs() {
	t.a = 10
	t.b = 10
}

func main() {
	t := Test{5, 5}

	t.Abs()

	fmt.Println(t)

	t.PointerAbs()

	fmt.Println(t)
}