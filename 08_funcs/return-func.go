package main

import "fmt"

func main() {

	x := retornafunc()

	y := x(3)

	fmt.Println(y)
	fmt.Println(retornafunc()(3))

}

func retornafunc() func(int) int {

	return func(i int) int {
		return i * 10
	}
}
