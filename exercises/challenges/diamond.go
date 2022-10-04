/*Monte um diamante a partir da logica dos loops*/

package main

import "fmt"

var m, n, number int

func main() {
	// number = 7

	fmt.Print("Digite o numero de linhas: ")
	fmt.Scan(&number)

	for m = 1; m <= number; m++ {

		for n = 1; n <= number-m; n++ {
			fmt.Print(" ")
		}
		for n = 1; n <= (m*2)-1; n++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for m = number - 1; m > 0; m-- {

		for n = 1; n <= number-m; n++ {
			fmt.Print(" ")
		}

		for n = 1; n <= (m*2)-1; n++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
