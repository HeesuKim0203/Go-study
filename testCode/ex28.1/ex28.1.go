package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	num := square(9)
	fmt.Printf("9 * 9 = %d\n", num)
}
