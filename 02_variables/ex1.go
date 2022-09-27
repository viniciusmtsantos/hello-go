// 2. Use fmt . Scanf to read a user's name and print " Hello < NAME > " with < NAME > replaced with the user's name to the terminal .

package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter with your name: ")
	fmt.Scanf(name)
	fmt.Println("Hello, ", name)
}
