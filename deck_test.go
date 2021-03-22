package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 32 {
		t.Errorf("Expected 32 cards, but got %v", len(cards))
	}

	if cards[0] != "Spade of One" {
		t.Errorf("Expected to be Spade of One, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Club of Eight" {
		t.Errorf("Expected to be Club of Eight, but got %v", cards[len(cards)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 32 {
		t.Errorf("Expected to be 32 cards, but got %v", len(loadedDeck))
	}
}
