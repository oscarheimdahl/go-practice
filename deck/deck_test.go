package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected test length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveAndLoadToFile(t *testing.T) {
	os.Remove("_decktesting")
	cards := newDeck()
	cards.saveToFile("_decktesting")
	cards2 := newDeckFromFile("_decktesting")

	if len(cards) != len(cards2) {
		t.Errorf("Expected deck to be the same length before and after saving and reading to disc, but loaded deck was of length: %v", len(cards2))
	}

	for i := range cards {
		if cards[i] != cards2[i] {
			t.Errorf("Expected read deck from disc to look the same as saved deck to disc")
		}
	}
}
