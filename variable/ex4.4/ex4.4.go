package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 20.0

	var c int = int(b)
	d := float64(a) * b

	var e int64 = int64(7)
	f := int64(a) * e

	fmt.Println(a, b, c, d, e, f)
}
