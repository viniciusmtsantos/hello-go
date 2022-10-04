package main

import "fmt"

// o struct nos permite armazenar dados com tipos diferentes
func main() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 6,
	}

	fmt.Println(x)

}
