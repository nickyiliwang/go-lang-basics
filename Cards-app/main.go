package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// slicing an array: cards[startIndexIncluding : endingToNotIncluding]

func main() {
	// cards := newDeck()
	// cards.print("Deck:")

	// hand, _ := deal(cards, 5)

	// hand.print("Hand:")
	// remainingDeck.print("Remain:")

	// hand.saveToFile("hand dealt")

	// reading deck from file
	// cards := newDeckFromFile("hand dealt")

	// Shuffling
	cards := newDeck()
	cards.shuffle()
	cards.print("Post Shuffle")

}

//deck.go ////////////////////////////////////////////////////

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

	// fmt.Println(cards)
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

// when we want to save info to disk as a file we need the deck type to be a string then to a byte
func (d deck) toString() string {
	// turn ["Ace of Clubs", "Two of Clubs","Three of Clubs"]
	// into "Ace of Clubs, Two of Clubs, Three of Clubs"
	// Type conversion: result := []typeYouWant(valueYouHave)
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// 0666 is default permission: anyone can read
	// returning only error values
	// ioutil is stn lib
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// byteslice and error, think of them as strings
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and reutrn a call to newDeck()
		// Option 2 - log the error and quit the program
		fmt.Println("Error:", err)
		// any int that's not 0 indicate an need to exit
		os.Exit(1)
		// Option 3 - PANIC!!!!
		// panic(err)
	}

	ss := strings.Split(string(bs), ", ") // spliting the byteslice converted string into string slice

	return deck(ss)
}

// no return value or params, just take the deck value and shuffle the values inside it
func (d deck) shuffle() {
	// creating none fixed seed
	seed := time.Now().UnixNano()
	// new rand type with new seed, now gens new randomness each time
	r := rand.New(rand.NewSource(seed))
	for i := range d {
		// stn pkg for generating random integer
		newPosition := r.Intn(len(d) - 1)
		// swapping the two positions in our desk slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
