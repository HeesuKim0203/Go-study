package main

import (
	"fmt"
)

type Test struct {
	X, Y int
}

func (v Test) Abs() int {
	return v.X + v.Y
}

func main() {
	v := Test{3, 4}
	fmt.Println(v.Abs())
}
