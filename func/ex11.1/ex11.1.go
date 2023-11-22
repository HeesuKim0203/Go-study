package main

import (
	"fmt"
)

func test(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	
	add := func (a int, b int) int {
		return a + b
	}
	
	fmt.Println(test(add))
}
