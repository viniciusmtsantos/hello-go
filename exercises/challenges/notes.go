/*Escreva um procedimento que recebe as 3 notas de um aluno por parâmetro e uma letra. Se a letra for A o procedimento calcula a média aritmética das notas do aluno, se for P, a sua média ponderada (pesos: 5, 3 e 2) e se for H, a sua média harmônica. A média calculada também deve retornar por parâmetro.
 */
package main

import "fmt"

var n1, n2, n3, m float32
var media, tipo string

func main() {
	fmt.Print("Digite a nota 1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite a nota 2: ")
	fmt.Scan(&n2)
	fmt.Print("Digite a nota 3: ")
	fmt.Scan(&n3)
	fmt.Println("Selecione:\nA-Aritmética\nP-Ponderada\nH-Harmônica")
	fmt.Scan(&media)
	notas(media)
}

func notas(tipo string) {
	if tipo == "A" || tipo == "a" {
		m = (n1 + n2 + n3) / 3
		fmt.Println("A média aritmética é: ", m)
	}
	if tipo == "P" || tipo == "p" {
		m = (n1*5 + n2*3 + n3*2) / 10
		fmt.Println("A média ponderada é: ", m)
	}
	if tipo == "H" || tipo == "h" {
		m = 0
		fmt.Println("nem sei como faz isso")
	}
}
