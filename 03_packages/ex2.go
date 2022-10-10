package main

import "fmt"

func main() {
	x := "Oi,"
	y := "Bom dia"

	// fmt.Print() // Print sem a linha nova no final
	// fmt.Println() // print que vem com uma linha nova no final
	y = fmt.Sprint(x, " ", y, "\n") // Ele retorna uma string e vai me dar um retorno desta função
	fmt.Print(y)

}

// func loop() {
// 	for i := 0; i <= 999; i++ {
// 		fmt.Println(i, " - ", string(i))
// 	}
// }
