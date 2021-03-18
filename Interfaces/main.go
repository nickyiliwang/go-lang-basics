package main

import "log"

type Champion interface {
	// defining a function that every champion must have
	SayName() string
	SayQuote() string
	NumberOfWeapons() int
}

type Rengar struct {
	Name    string
	Role    string
	Weapons int
}

type Soraka struct {
	Name     string
	Role     string
	Position string
	Weapons  int
}

func main() {
	rengar := Rengar{
		Name:    "Rengar",
		Role:    "Assassin",
		Weapons: 2,
	}

	PrintMyChamp(rengar)

	soraka := Soraka{
		Name:     "Soraka",
		Role:     "Support",
		Position: "Bottom",
		Weapons:  1,
	}

	// Error when passing soraka type into PrintMyChamp, using the Champion interface, because soraka did not have the required signature, aka required func.
	PrintMyChamp(soraka)
}

// we created 2 functions and
// (receiver type) (no params)
func (r Rengar) SayName() string {
	return r.Name
}

func (r Rengar) SayQuote() string {
	return r.Name + " the true hunter never rests"
}

func (r Rengar) NumberOfWeapons() int {
	return r.Weapons
}

func (s Soraka) SayName() string {
	return s.Name
}

func (s Soraka) SayQuote() string {
	return s.Name + " will not heal you!"
}

func (s Soraka) NumberOfWeapons() int {
	return s.Weapons
}

func PrintMyChamp(c Champion) {
	log.Println(c.SayName(), "says:", c.SayQuote())
	log.Println(c.SayName(), "uses", c.NumberOfWeapons(), "weapons to fight")
}
