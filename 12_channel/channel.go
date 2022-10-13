package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
	}()

	// result := <-canal

	fmt.Println(<-canal) // se não comentar ele passa o bastão
	// fmt.Println("O result não recebe o valor do canal pq ele ja foi deixado acima", result)

}
