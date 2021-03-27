package main

import (
	"log"
	"time"
)

type Champion struct {
	Name         string
	Class        string
	AttackDamage int
	AbilityPower int
	ReleaseDate  time.Time
}

func (name *Champion) printChampionName() string {
	return name.Name
}

func main() {
	var rengar Champion

	rengar.Name = "Rengar"
	rengar.AttackDamage = 100

	nunu := Champion{
		Name: "nunu",
	}

	// Using a pointer to print champ name
	log.Println("first champ is:", rengar.printChampionName())
	log.Println("second champ is:", nunu)
}
