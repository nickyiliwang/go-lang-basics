package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// no receiver in this fn because we want to return a new deck instance
func newDeck() deck {
	// cards := deck{}

	// cardSuits := 
}

// receiver
// the receiver can be thought of as the "this" key word in js
// d is reffering to the actual copy of the deck, in this case: cards
// deck: every variable of type 'deck' can now call this func on itself
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
