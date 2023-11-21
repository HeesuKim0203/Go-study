package main

import "fmt"

func main() {
	a := 2
	sum := 1
	for ; sum < 10 ; {
		fmt.Println(a, " + ", sum, " = ", a * sum)
		sum++
	}
}
