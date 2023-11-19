package main

import "fmt"

func main() {
	var a int = 3
	var b int
	c := 10

	fmt.Printf("%d + %d = %d\n", a, b, a+b)

	fmt.Printf("%d + %d = %d", c, a+b, a+b+c)
}
