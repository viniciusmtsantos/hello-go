package main

import "fmt"

func main() {
	/*- Qual a utilidade de funções ?
	      - Abstrair funcionalidade
	      - Reutilização de código
	  - func ( receiver ) identifier
	  ( parameters ) ( returns ) { code }
	  - A diferença entre parâmetros e
	  argumentos :
	  - Funções são definidas com parâmetros
	  - Funções são chamadas com argumentos
	  - Tudo em Go é * pass by value . *
	         Pass by reference , pass by
	      copy , ... não .*/

	/*
		s := "abc"
		fmt.Println(len(s)) // len é uma função que retorna o numero de caracteres dentro da variável
	*/

	basic()
	argumento("tarde")

}

func basic() {
	fmt.Println("Olá, bom dia")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Olá, Bom dia")
	} else if s == "tarde" {
		fmt.Println("Olá, Bom tarde")
	} else {
		fmt.Println("Noite")
	}
}
