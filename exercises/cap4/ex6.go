// Utilizando iota , crie 4 constantes cujos valores sejam os próximos 4 anos .

package main

import "fmt"

const (
	ano1 = 2022
	ano2 = ano1 + iota
	ano3
	ano4
	ano5
	x int = 42
)

func main() {
	fmt.Printf("%x", x)
	fmt.Printf("Ano atual é: %v\n", ano1)
	fmt.Println(ano2)
	fmt.Println(ano3)
	fmt.Println(ano4)
	fmt.Println(ano5)

}
