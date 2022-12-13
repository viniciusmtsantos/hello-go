// Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

func main() {
	x := 10
	for x <= 100 {
		rest := x % 4
		fmt.Println(rest)
		x++
	}
}
