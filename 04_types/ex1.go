package main

import "fmt"

type hotdog int // declarando o tipo da variável

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // vemos o tipo da variável x
	x = 42
	fmt.Println(x)
}
