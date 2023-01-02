/*- Crie uma slice usando make que
possa conter todos os estados do
Spas aa
- Os estados: "Acre", "Alagoas",
"Amapa", "Amazonas", "Bahia",
"Ceara", "Espirito Santo",
"Goids", "Maranhao", "Mato
Grosso", "Mato Grosso do Sul",
"Minas Gerais", "Para", “Paraiba”,
"Parana", "Pernambuco", "Piaui",
"Rio de Janeiro", "Rio Grande do
Norte", "Rio Grande do Sul",
"Rond6énia", "Roraima", "Santa
Catarina", "Sao Paulo", "Sergipe",
"Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice




*/package main

import "fmt"

func main() {
	estados := make([]string, 26, 26) // 5 elementos e 10 de capacidade

	estados = []string{"Acre", "Alagoas",
		"Amapa", "Amazonas", "Bahia",
		"Ceara", "Espirito Santo",
		"Goids", "Maranhao", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Para", "Paraiba",
		"Parana", "Pernambuco", "Piaui", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
		"Rond6énia", "Roraima", "Santa Catarina", "Sao Paulo", "Sergipe",
		"Tocantins"}
	fmt.Println(len(estados), cap(estados))
}
