package main

import "fmt"

func main() {
	esportefavorito := "voley"
	switch true {
	case esportefavorito == "futebol":
		fmt.Printf("Voce curte %v", esportefavorito)
	case esportefavorito == "voley":
		fmt.Printf("Voce curte %v", esportefavorito)
	case esportefavorito == "Basket":
		fmt.Printf("Voce curte %v", esportefavorito)
	default:
		fmt.Println("Não gosta de nem de voley, nem de futebol, nem de basket ")
	}

}
