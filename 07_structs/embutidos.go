package main

import "fmt"

// o struct nos permite armazenar dados com tipos diferentes

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Marcos",
			idade: 31,
		},
		titulo:  "Dev em Go",
		salario: 100000,
	}

	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2.titulo)
}
