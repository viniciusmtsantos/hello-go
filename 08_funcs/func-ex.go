package main

import "fmt"

func main() {

	x := 123

	y := func(x int) int {
		// fmt.Println(x, " vezes 87354 é: ")
		return x * 87354
	}
	fmt.Println(x, "vezes 87354 é:", y(x))
}
