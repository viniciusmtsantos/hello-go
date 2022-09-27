package main

import "fmt"

const (
	x = iota + 1000 // define valores para as constantes sendo uma diferente da outra
	y = iota
	z = iota
)

func main() {
	fmt.Println(x, y, z)
}
