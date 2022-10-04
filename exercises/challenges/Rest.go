/*
Faça um programa que verifique e mostre os números entre 1.000 e 2.000 (inclusive) que, quando divididos por 11 produzam resto igual a 2.
*/
package main

import "fmt"

func main() {
	for a := 1000; a <= 2000; a++ {
		if a%11 == 2 {
			fmt.Print(a, ", ")
		} else {
		}
	}
	fmt.Println("FIM")
}
