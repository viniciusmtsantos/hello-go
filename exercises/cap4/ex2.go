package main

import "fmt"

var (
	x bool
)

func main() {
	x = 4 > 5
	fmt.Println(x)
	y := (false == x)
	fmt.Println(y)
}
