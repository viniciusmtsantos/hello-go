package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Enter the First number: ")
	fmt.Scan(&x)
	fmt.Print("Enter the Second number: ")
	fmt.Scan(&y)
	for x < y {
		x++
		fmt.Println(x)
	}

}
