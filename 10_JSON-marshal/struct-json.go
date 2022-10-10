package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	vinicius := Pessoa{"vinicius", 2}
	j, err := json.Marshal(vinicius)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))

}
