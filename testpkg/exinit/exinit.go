package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init() d : ", d)
}

func f() int {
	d++
	fmt.Println("f() d : ", d)
	return d
}

func PrintD() {
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
}
