/*
- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um ponteiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
*/

package main

import "fmt"

func main() {
	x := 10

	y := &x
	fmt.Println(*&x) // conteúdo do endereço da variável x
	fmt.Println(*y)  // fazer um dereference

}
