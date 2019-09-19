package main

import (
	"fmt"
)

type my_number int

var x my_number

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
