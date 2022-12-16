/* Usando uma Literal composta:
- Crie um array que suporte 5 valores do tipo int
- Atribua valores aos seus indices
Utilize range e demonstre os valores do array.

Utilizando format printing, demonstre o tipo do array.
*/

package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	for i, v := range array {
		fmt.Println(i, ":", v)
	}
	fmt.Printf("%T", array)
}
