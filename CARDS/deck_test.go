package main

import "testing"

//captialize function name to export
func TestNewDeckLength(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Aced of spades, but got %v", d[0])
	}

}

func TestLastCard(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expeceted last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}
