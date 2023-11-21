package main

import "fmt"

func mulTest(a, b int) int {

	if v := 5 ; a < 0 {
		return v * b 
	}

	return a * b
} 

func main() {
	fmt.Println(
		mulTest(3, 2),
		mulTest(-3, 3),
	)
}
