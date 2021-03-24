package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// no receiver in this fn because we want to return a new deck instance
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			card := val + " of " + suit
			cards = append(cards, card)
		}
	}

	fmt.Println(cards)
	return cards
}

// receiver
// the receiver can be thought of as the "this" key word in js
// d is reffering to the actual copy of the deck, in this case: cards
// deck: every variable of type 'deck' can now call this func on itself
func (d deck) print(msg string) {
	fmt.Println(msg)
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
