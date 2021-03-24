package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	// returns a new slice back to the variable of cards
	cards = append(cards, "Six of Diamonds")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
