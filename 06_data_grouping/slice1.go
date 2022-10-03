package main

import "fmt"

func main() {

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12} // estes numeros são uma fatia de ints

	umaslice = append(umaslice, 5, 6, 7, 8)    // são ints
	umaslice = append(umaslice, outraslice...) // fatia de indices
	fmt.Println(umaslice)

}
