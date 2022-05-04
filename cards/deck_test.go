package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveDeckAndReadDeck(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveDeck("_decktesting")
	loadedDeck := readDeck("_deckTesting")

	if deck.toString() != loadedDeck.toString() {
		t.Errorf("Expected newDeck: %v \nbut got \n%v \n instead", deck, loadedDeck)
	}

	os.Remove("_decktesting")
}
