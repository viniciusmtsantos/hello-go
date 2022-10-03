package main

import (
	"fmt"
)

func main() {

	// amigos := map[string]int{
	// 	"alfredo": 5551234,
	// 	"joana":   5674442,
	// }
	// amigos["gopher"] = 44444

	// // comma ok idiom

	// fmt.Println(amigos)
	// fmt.Println(amigos["joana"])
	// if será, ok := amigos["fantasma"]; !ok {
	// 	fmt.Println("Não existe")
	// } else {
	// 	fmt.Println(será)
	// }

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "Esse é massa",
		18:  "idade de ir",
	}
	fmt.Println(qualquercoisa)

	for key, value := range qualquercoisa {
		fmt.Println(key, ":", value)
	}

	// deletando um elemento de um map
	delete(qualquercoisa, 123)
}
