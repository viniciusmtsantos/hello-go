/* Crie um loop utilizando a sintaxe : for condition { }
- Utilize - o para demonstrar os anos desde que você nasceu .
- Solução : https://play.golang.org/p/PzUqGLODmW
*/

package main

import "fmt"

func main() {
	birth := 1999
	atual := 2022
	for year := atual; year >= birth; year-- {
		fmt.Println(year)
	}
}
