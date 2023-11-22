package main

import "fmt"

func main() {
	s := []int{ 2, 10, 2, 18 }
	
	s = append(s, 2)

	fmt.Println(s)
	
	for a, b := range(s) {
		fmt.Println(a, b)
	}
}
