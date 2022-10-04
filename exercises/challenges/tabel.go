/*Faça um programa que receba um número e usando laços de repetição calcule e mostre a tabuada desse número.
 */
package main

import "fmt"

func main() {
	var numOne int
	fmt.Print("Enter a number: ")
	fmt.Scan(&numOne)
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v * %v = %v\n", numOne, i, numOne*i)

	}
}
