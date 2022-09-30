package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pokemon1Attack string
var pokemon2Attack string
var index int
var x, y int

type pokemon struct {
	name  string
	types string
	moves []string
	stats map[string]int
	hp    int
}

var pokemon1 = pokemon{
	name:  "Charizard",
	types: "Fire",
	moves: []string{"Flamethrower", "Fly", "Blast Burn", "Fire Punch"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    100,
}

var pokemon2 = pokemon{
	name:  "Blastoise",
	types: "Water",
	moves: []string{"Water Gun", "Bubblebeam", "Hydro Pump", "Surf"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    100,
}

var pokemon3 = pokemon{
	name:  "Acrobatics",
	types: "Flying",
	moves: []string{"Magic Coat", "Power Fly", "Fire Fly", "Monster Fly"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    100,
}

var pokemon4 = pokemon{
	name:  "Aura",
	types: "Electricity",
	moves: []string{"Dark Atack", "Fly Electric", "Pain Max", "Water Electrified"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    100,
}

func main() {

	fight(pokemon1, pokemon2)
}

// print one character at a time
func delayPrint(s ...string) {
	for _, c := range s {
		fmt.Print(string(c))
		time.Sleep(180 * time.Millisecond)
	}
}

func (p pokemon) lessthanzero() {
	if p.hp < 0 {
		for p.hp != 0 {
			p.hp = p.hp + 1
		}
	}
}

func fight(one pokemon, two pokemon) {
	pokemons := []pokemon{}
	pokemons = append(pokemons, pokemon1, pokemon2, pokemon3, pokemon4)
	fmt.Println("-----POKEMON BATTLE-----")
	fmt.Print("0 - Charizard\n1 - Blastoise\n2 - Acrobatics\n3 - Aura\n\n")
	fmt.Print("Pick First Pokemon: ")
	fmt.Scanln(&x)
	for x > 3 {
		fmt.Print("Invalid Number!\n\nPick First Pokemon: ")
		fmt.Scanln(&x)
	}
	fmt.Println("First READY!")
	fmt.Println("NAME: ", pokemons[x].name)
	fmt.Println("TYPE/", pokemons[x].types)
	fmt.Println("ATTACK/", pokemons[x].stats["ATTACK"])
	fmt.Println("DEFENSE/", pokemons[x].stats["DEFENSE"])
	fmt.Print("\n-- VS --\n\n")
	fmt.Print("Pick Second Pokemon: ")
	fmt.Scanln(&y)
	for y > 3 {
		fmt.Print("Invalid Number!\n\nPick Second Pokemon: ")
		fmt.Scanln(&y)
	}
	fmt.Println("Second READY!")
	fmt.Println("NAME: ", pokemons[y].name)
	fmt.Println("TYPE/", pokemons[y].types)
	fmt.Println("ATTACK/", pokemons[y].stats["ATTACK"])
	fmt.Println("DEFENSE/", pokemons[y].stats["DEFENSE"])
	time.Sleep(2 * time.Second)

	// consider types advantages
	version := []string{"Fire", "Water", "Flying", "Electricity"}

	for i, k := range version {
		if pokemons[x].types == k {
			// both are same types
			if pokemons[y].types == k || pokemons[x].types == k {
				pokemon1Attack = "\nIts not very effective...\n"
				pokemon2Attack = "\nIts not very effective...\n"
			}

			// pokemon2 is STRONG
			if pokemons[y].types == version[(i+1)%3] {
				pokemons[y].stats["ATTACK"] *= 2
				pokemons[y].stats["DEFENSE"] *= 2
				pokemons[x].stats["ATTACK"] /= 2
				pokemons[x].stats["DEFENSE"] /= 2
				pokemon1Attack = "\nIts not very effective...\n"
				pokemon2Attack = "\nIts super effective!\n"
			}

			// pokemon2 is WEAK
			if pokemons[y].types == version[(i+2)%3] {
				pokemons[x].stats["ATTACK"] *= 2
				pokemons[x].stats["DEFENSE"] *= 2
				pokemons[y].stats["ATTACK"] /= 2
				pokemons[y].stats["DEFENSE"] /= 2
				pokemon1Attack = "\nIts super effective!\n"
				pokemon2Attack = "\nIts not very effective...\n"
			}
		}

	}

	// Actual fight
	// Continue while pokemon still have health
	for {

		// begins battle
		fmt.Printf("\n%v \t\tHLTH:\t%v\n", pokemons[x].name, pokemons[x].hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemons[y].name, pokemons[y].hp)

		// show moves
		fmt.Printf("Go %v!\n\n", pokemons[x].name)
		for i, x := range pokemons[x].moves {
			fmt.Printf("%v. %v\n", i+1, x)
		}

		// first turn battle
		fmt.Print("Pick a move: ")
		fmt.Scanln(&index)
		delayPrint("\n", pokemons[x].name, " used ", pokemons[x].moves[index-1], "!")
		time.Sleep(1 * time.Second)
		delayPrint(pokemon1Attack)

		// determine damage
		pokemons[y].hp -= pokemons[x].stats["ATTACK"]

		// adiciona mais 1 ao hp até ficar com 0 para que não termine com o hp negativo
		pokemons[y].lessthanzero()
		// fim da implementação

		time.Sleep(1 * time.Second)
		fmt.Printf("\n%v \t\tHLTH:\t%v\n", pokemons[x].name, pokemons[x].hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemons[y].name, pokemons[y].hp)
		time.Sleep(1 * time.Second)

		// check if pokemon fainted
		if pokemons[y].hp <= 0 {
			delayPrint("\n...", pokemons[y].name, " fainted.")
			break
		}

		// pokemon2 turn
		// show moves
		fmt.Printf("Go %v!\n", pokemons[y].name)
		for i, x := range pokemons[y].moves {
			fmt.Printf("%v. %v\n", i+1, x)
		}

		// second turn battle
		fmt.Print("Pick a move: ")
		fmt.Scanln(&index)
		delayPrint("\n", pokemons[y].name, " used ", pokemons[y].moves[index-1], "!")
		time.Sleep(1 * time.Second)
		delayPrint(pokemon2Attack)

		// determine damage
		pokemons[x].hp -= pokemons[y].stats["ATTACK"]

		// adiciona mais 1 ao hp até ficar com 0 para que não termine com o hp negativo
		pokemon1.lessthanzero()

		time.Sleep(1 * time.Second)
		fmt.Printf("\n%v \t\tHLTH:\t%v\n", pokemons[x].name, pokemons[x].hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemons[y].name, pokemons[y].hp)
		time.Sleep(1 * time.Second)

		// check if pokemon1 fainted
		if pokemons[x].hp <= 0 {

			delayPrint("\n...", pokemons[x].name, " fainted.")
			break
		}

		if pokemons[x].hp > 0 && pokemons[y].hp > 0 {
			continue
		}

	}
	rand.Seed(time.Now().UnixNano())
	money := rand.Intn(3000)
	delayPrint("\nOpponent paid you: ")
	fmt.Printf("%vG\n", money)

}
