/*- WP : " The most common application of recursion is in mathematics and computer science , where a
function being defined is applied within its own definition . "
- Exemplos de recursividade : Fractais , matrioscas , efeito Droste ( o efeito produzido por uma imagem
que aparece dentro dela própria ) , GNU ( " GNU is Not Unix " ) , etc.
- No estudo de funções : é uma função que chama a ela própria .
- Exemplo : fatoriais .
      4 ! = 4 * 3 * 2 * 1 ( e no zero , deu . )
*/
package main

import "fmt"

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(loops(6))

}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
