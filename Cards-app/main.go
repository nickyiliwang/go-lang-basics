package main

// slicing an array: cards[startIndexIncluding : endingToNotIncluding]

func main() {
	// cards := newDeck()
	// cards.print("Deck:")

	// hand, _ := deal(cards, 5)

	// hand.print("Hand:")
	// remainingDeck.print("Remain:")

	// hand.saveToFile("hand dealt")

	// // reading deck from file
	// cards := newDeckFromFile("hand dealt")

	// Shuffling
	cards := newDeck()
	cards.shuffle()
	cards.print("Post Shuffle")

}
