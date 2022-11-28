// package main

// /*
//   Go é modular e usa o conceito de pacotes para
//   organizar o sistema e ajudar na reutilização do
//   código do mesmo.
//   Então a primeira coisa a se fazer quando
//   criamos um arquivo .go e declarar o pacote do
//   qual ele faz parte.
//   No caso desse arquivo "package main". É uma KEYWORD, uma palavra chave, uma identificadora
// */

// import "fmt"

// /*
//   import é utilizado para chamar outros pacotes
//   da linguagem, de terceiros e pacotes
//   que você tenha criado.
//   Nesse caso "fmt" é um pacote para formatação
//   de dados de saída.
// */
// var y = "Ola, Bom dia" // Aqui precisa do var

// func main() {
// 	// o x e o z funcionaram apenas dentro do scope
// 	x := 10 + 10                              // 10 + 10 é uma expressão -> algo que produz um resultado, mas não gera uma ação ; := é o Operador curto de declaração -> TIPAGEM AUTOMÁTICA Estou declarando uma variável e atribuindo um tipo a ela automaticamente
// 	z := 10.0                                 // A marmota tem que ser usada obrigatoriamente dentro do codeblock
// 	fmt.Printf("x: %v, é um tipo %T\n", x, x) // uma statement é uma linha de código que forma uma ação, neste caso uma chamada de função
// 	fmt.Printf("y: %v, é um tipo %T\n", y, y)
// 	fmt.Printf("z: %v, é um tipo %T\n", z, z)

// 	fmt.Println("Hello, World!")

// }

// /*ss
//   main() é a função principal de um arquivo ".go".
//   Ela é a espinha dorsal do programa.
//   Através dela que as demais funções,declarações
//   de variáveis, tipo e métodos serão executados.
// */

// // execute-o como o comando: go run intro.go. A mensagem “Hello World” deve ser exibida no console do seu terminal.
