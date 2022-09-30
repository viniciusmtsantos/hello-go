package main

import "fmt"

func main() {
	var numOne, numTwo, remainder int
	fmt.Print("Enter a number one: ")
	fmt.Scan(&numOne)
	fmt.Print("Enter a second number one: ")
	fmt.Scan(&numTwo)
	remainder = numOne % numTwo
	fmt.Print("The remainder result is: ", remainder)
}
