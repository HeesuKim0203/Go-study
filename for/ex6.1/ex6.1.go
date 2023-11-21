package main

import "fmt"

func main() {
	var a int = 5 ;
	
	for i := 1 ; i < 10 ; i++ {
		fmt.Println(a, " * ", i, " = ", a * i)
	}
}
