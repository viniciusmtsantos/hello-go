// - Crie um loop utilizando a sintaxe : for { } Utilize - o para demonstrar os anos desde que você nasceu

package main

import "fmt"

func main() {
	birth := 1999
	atual := 2022
	for {
		fmt.Println(birth)
		birth++
		if birth > atual {
			break
		}
	}
}
