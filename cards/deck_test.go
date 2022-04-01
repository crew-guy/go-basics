package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16,but got %v", len(d))
	}

	if d[0] != "One of Hearts" {
		t.Errorf("Expected 1st element to be One of Hearts, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 1st element to be One of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of loaded deck to be 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
