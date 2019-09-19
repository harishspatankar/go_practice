package main

import (
	"fmt"
)

type my_number int

var x my_number
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
