package main

import "fmt"

var x int = 42
var y string = "James"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z) // SprintF é usado para atribuir todos os valores a uma única variável, neste caso s
	fmt.Println(s)
}
