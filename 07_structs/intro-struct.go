package main

import "fmt"

// o struct nos permite armazenar dados com tipos diferentes

type x struct {
	nome      string
	sobrenome string
	casado    bool
}

func main() {
	x1 := x{
		nome:      "João",
		sobrenome: "Uou",
		casado:    false,
	}
	x2 := x{"Vinicius", "Santos", true}

	fmt.Println(x1, x2)

}
