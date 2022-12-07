/*
Põe na tela : 0 unicode code point de todas as letras maiúsculas do alfabeto , três vezes cada .
- Por exemplo :
  65
  U + 0041 ' A '
  U + 0041 ' A '
  U + 0041 ' A '
  66
  U + 0042 ' B '
  U + 0042 ' B '
  U + 0042 ' B '
  ... e por aí vai .

*/
package main

import "fmt"

func main() {
	for letters := 65; letters <= 90; letters++ {
		fmt.Println(letters)
		for n := 1; n <= 3; n++ {
			fmt.Printf("\t%#U\n", letters)
		}
	}
}
