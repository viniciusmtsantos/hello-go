package main

import "fmt"

// Podemos declarar variáveis de diversas formas agrupadas usando var (...), ou individualmente usando var string , ou simplesmente atribuindo um alor v:="Teste
// Observe abaixo:

// Em grupo
// Go tem é uma ling estaticamente tipada, isto é, a definição do tipo da variável fora do code block define permanentemente o tipo da variável
var (
	//determinando um tipo e atribuindo um valor
	nome string = "Vinícius"
	// determinando um tipo, sem atribuição de valor
	cidade string
	idade  int
	casado bool
	y      = 6
)

func main() {
	fmt.Printf("%v -> %T \n", cidade, cidade)
	fmt.Printf("%v -> %T \n", idade, idade)
	fmt.Printf("%v -> %T \n", casado, casado)

	idade = 23
	casado = true
	// "carteira" não foi declarada previamente. Nesse caso ao atribuirmos um valor usando ":=" a variável passa a ser do tipo atribuído "carteira := 15.35"
	carteira := 00.15
	// O mesmo ocorre com "estado_civil" que agora é do tipo string
	estado_civil := "casado com Ludmilla"

	// Mais para frente falaremos sobre estruturas de controle então não se incomode com esse bloco agora
	fmt.Println(nome, "tem", idade, "anos de idade e é", estado_civil)
	fmt.Println(nome, "tem R$", carteira, " na carteira")
	qualquercoisa(y)
}

func qualquercoisa(x int) {
	fmt.Println(x)
}
