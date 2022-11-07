/* Crie um programa que :
   Atribua um valor int a uma variável
  - Demonstre este valor em decimal , binário e hexadecimal
  - Desloque os bits dessa variável 1 para a esquerda , e atribua o resultado a outra variável
	Demonstre esta outra variável em decimal , binário e hexadecimal
*/
package main

import (
	"fmt"
)

const x = iota + 1

func main() {
	fmt.Println("Para X:")
	fmt.Printf("O numero Decimal é: %d\n", x)
	fmt.Printf("Em Binário: %b\n", x)
	fmt.Printf("Em Hexadecimal: %#x\n", x)
	// y := strconv.FormatInt(int64(x), 2) // convertendo em binário
	y := x << 1
	fmt.Println("Para Y:")

	fmt.Printf("O numero decimal é: %d\n", y)
	fmt.Printf("Em Binário: %b\n", y)
	fmt.Printf("Em Hexadecimal: %#x\n", y)

}
