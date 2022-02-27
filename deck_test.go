package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	if d[51] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v", d[0])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(loadedDeck))
	}

	if loadedDeck[0] != deck[0] {
		t.Errorf("Expected first card to be %v, but got %v", deck[0], loadedDeck[0])
	}

	if loadedDeck[51] != deck[51] {
		t.Errorf("Expected last card to be %v, but got %v", deck[51], loadedDeck[51])
	}

	os.Remove("_decktesting")

}
