package main

import "fmt"

// array do tipo cinco int
var x [5]int

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)
	fmt.Println(slice2[3])

}
