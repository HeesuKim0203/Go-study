package main

import (
	"fmt"

	"src/testpkg/custompkg"
	"src/testpkg/exinit"

	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	fmt.Println("Test Module and Package")
	custompkg.PrintCustom()

	expkg.PrintSample()

	exinit.PrintD()
}
