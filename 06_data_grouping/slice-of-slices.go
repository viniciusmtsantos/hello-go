package main

import "fmt"

// array do tipo cinco int

func main() {

	/* slices são tipos de dados compostos
	- Função range

	slice := []string{"banana", "maça", "laranja"}
	fmt.Println(slice)
	// slice[3] = "melancia"
	slice = append(slice, "melancia", "pera")
	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor: ", valor)
	}
	// for _, valor := range slice {
	// 	fmt.Println("Um dos valores é: ", valor)

	// }
	for indice, _ := range slice {
		fmt.Println("Os indices são ", indice)
	}
	slice2 := []int{23, 45, 36}
	total := 0

	for _, value := range slice2 {
		total += value
		fmt.Println(total)
	}
	*/

	/*slice de slices
	//                     0             1            2         3                 4
	sabores := []string{"peperoni", "muzzarela", "abacaxi", "quatroquejo", "marguerita"}
	// fatia := sabores[1:3]
	fatia_total := sabores[:]
	fmt.Println(fatia_total)
	fmt.Println(sabores[0], sabores[3])

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	sabores = append(sabores[0:2], sabores[3:]...)
	fmt.Println(sabores)
	*/

	ss := [][]int{

		// Indice: 0 1 2       // Indice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4])
}
