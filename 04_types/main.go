package main

import (
	"fmt"
)

type hotdog int // criando um type chamado hotdog e o tipo subjacente dele é o int

var (
	x int
	b hotdog = 101
)

func main() {
	x = 50
	fmt.Printf("%v -> %T\n", b, b)
	fmt.Printf("%v -> %T\n", x, x)

	x = int(b)
	fmt.Printf("%v -> %T\n", x, x)

}
