/* - Utilizando como base o exercicio anterior, utilize slicing para
demonstrar os valores:
- Do primeiro ao terceiro item do slice (incluindo o terceiro
item!)
- Do quinto ao ultimo item do slice (incluindo o dltimo item!)
- Do segundo ao sétimo item do slice (incluindo o sétimo item! )
- Do terceiro ao penúltimo item do slice (incluindo o penúltimo
item!)
- Desafio: obtenha o mesmo resultado acima utilizando a funcao len() para determinar o penúltimo item
*/

package main

import "fmt"

func main() {
	// posição sempre começa do 0
	firstslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	secondslice := append(firstslice[0:3])
	fmt.Println(secondslice)

	terslice := append(firstslice[4:]) // o quinto item esta na posição 4
	fmt.Println(terslice)

	quadslice := append(firstslice[1:7]) // o setimo item esta na posição 6, mas com a fatia pega o proximo
	fmt.Println(quadslice)

	quintslice := append(firstslice[2:9]) // o setimo item esta na posição 8, mas com a fatia pega o proximo
	fmt.Println(quintslice)

}
