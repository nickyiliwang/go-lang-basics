package main

// slicing an array: cards[startIndexIncluding : endingToNotIncluding]

func main() {
	cards := newDeck()
	cards.print("Deck:")

	hand, remainingDeck := deal(cards, 5)

	// fmt.Println("hand:" + hand.print())
	// fmt.Println("remains:" + remainingDeck.print())
	hand.print("Hand:")
	remainingDeck.print("Remain:")
}
