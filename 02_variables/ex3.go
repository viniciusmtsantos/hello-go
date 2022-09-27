package main

import "fmt"

type hotdog int // declarando o tipo da variável que tem um tipo subjacente int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // vemos o tipo da variável x
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println("Para baixo é y e para cima é x")
	fmt.Printf("%T\n", y) // vemos o tipo da variável x
	fmt.Println(y)
}
