package main

import "fmt"

func main() {

	x := 0
	recebevalor(x)
	// recebeponteiro(&x)
	fmt.Println("Fora da função é: ", x)

}

func recebevalor(x int) {
	x++
	fmt.Println("Na função é:", x)
}

func recebeponteiro(x *int) {
	*x++
	fmt.Println("Na função defer, x é: ", *x)
}
