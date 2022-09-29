package main

import "fmt"

func main() {

	quemtanoescritoriohoje := " ninguem "

	switch quemtanoescritoriohoje {
	case " zezinho ":
		fmt.Println(" hoje quem tá no escritório é o zezinho ")
		fallthrough
	case " marquinhos ":
		fmt.Println(" sempre  que o zezinho vem , o marquinhos vem tambem ")
	case " joana ":
		fmt.Println(" hoje quem tá no escritório é a joana ")
		fallthrough
	case " maria ":
		fmt.Println(" hoje quem tá no escritório é a maria ")
	default:
		fmt.Println(" tá vazio . ou a balada tava ótima , ou é feriado ")
	}
}

/*
	// for usual
	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}
	// declarando uma variável antes do for
	y := 0
	for y < 2 {
		y++
		fmt.Print(y)
	}
	// Sobre o Break
	z := 0
	for {
		if z < 10 {
			z++
			fmt.Println(" z é menor que dez ")
		} else {
			fmt.Println(" z é maior que dez mano tô fora ! ")
			break // o statement break para o loop teoricamente infinito
		}
	}
	// For com o continue faz com que a ação seja parada apenas para aquela situação, diferente do break que irá parar o FOR de fato
	for i := 0; i < 20; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
*/
/*
		 Switch :
		- pode avaliar uma expressão
		- switch statement == case ( value )
		- default switch statement == true ( bool )
		- não há fall - through por padrão
		- criando fall - through
		- default
		- cases compostos

	h := 6
	switch {
	case h > 5:
		fmt.Println("x é maior que 5")
	case h > 6:
		fmt.Println("x é maior que 6")
	case h > 8:
		fmt.Println("x é maior que 8")
	}
	x : = 8
	switch true {
	case ( x == 4 ) , ( x == 8 ) :
		fmt.Println( " é 4 ou 8 " )
		fallthrough
	case ( x < 10 ) :
		fmt.Println ( " 3 ou 4 " )
		default :
		// isso é o padrão se não houver outra opção
	}

*/
