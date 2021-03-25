package main

import "testing"

// t is a test handler, tell it what's wrong
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// 4 suits and 4 card values = 16 cards in total
	if len(d) != 16 {
		// if doesn't have 16 cards
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}

	// the first card should be ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	// the first card should be ace of spades
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}

}
