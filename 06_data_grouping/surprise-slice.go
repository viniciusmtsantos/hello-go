package main

import "fmt"

// array do tipo cinco int

func main() {

	firstslice := []int{1, 2, 3, 4, 5}
	secondslice := append(firstslice[:2], firstslice[4:]...)

	fmt.Println(secondslice)
	fmt.Println(firstslice)

}
