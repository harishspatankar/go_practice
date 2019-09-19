package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "Bond: James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%d\t %s\t %t\n", x, y, z)
}
