package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First element to be 'Ace of Spades' got: %v", d[0])
	}
	if d[len(d)-1] != "Jack of Diamonds" {
		t.Errorf("Expected Last element to be 'Jack of Diamonds' got: %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := deckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 go %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
