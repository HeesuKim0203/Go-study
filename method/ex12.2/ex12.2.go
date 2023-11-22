package main

import "fmt"

type Test int

func (t Test) Abs() int {
	return int(t)
}

func main() {
	t := Test(10)
	
	fmt.Println(t.Abs())
}