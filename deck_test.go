package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "T♠" {
		t.Errorf("Expected first card of T♠, but got %v", d[0])
	}

	if d[len(d)-1] != "4♣" {
		t.Errorf("Expected first card of 4♣, but got %v", d[0])
	}
}
