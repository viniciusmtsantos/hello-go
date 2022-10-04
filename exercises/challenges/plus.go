/*
Faça um programa que leia um valor n, inteiro e positivo, calcule e mostre a seguinte soma: S = 1 + 1/2 + 1/3 + 1/4 + … + 1/n.
*/

package main

import "fmt"

var number, soma, i float32

func main() {

	fmt.Print("Digite o numero: ")
	fmt.Scan(&number)
	for i = 1; i <= number; i++ {
		x := (1 / i)
		soma += x
		if i < number {
			fmt.Printf("(1/%v) + ", i)
		} else {
			fmt.Printf("(1/%v) = ", i)
		}
	}
	fmt.Println(soma)
}
