package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck :=  newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first deck value of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck) - 1] != "Four of Clubs" {
		t.Errorf("Expected last deck value of Four of Clubs, but got %v", deck[len(deck) - 1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_deckfiletesting"
	os.Remove(filename)

	deck := newDeck()

	deck.saveToFileName(filename)

	deckFromFile := newDeckFromFile(filename)

	if deck.toString() != deckFromFile.toString() {
		t.Errorf("Expected the same value of new deck %v, but got %v", deck, deckFromFile)
	}

	os.Remove(filename)
}