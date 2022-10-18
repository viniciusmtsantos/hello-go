package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndereco)
}
