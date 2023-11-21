package main

import "fmt"

func main() {
	var a int
	var b int
	var tmp int

	fmt.Scanln(&a) 
	fmt.Scanln(&b)

	tmp = b % 100

	fmt.Scanln()
}