package main

import "fmt"

func fibonacci() func() int {
	count := 0
	a := 0
	b := 1
	return func() int {
		sum := 0
		if count == 1 {
			sum = a + b
		} else if count > 1 {
			sum = a + b
			a = b
			b = sum
		}
		count++
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
