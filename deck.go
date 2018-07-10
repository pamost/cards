package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♠", "♦", "♥", "♣"}
	cardVlues := []string{"Туз", "2", "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardVlues {
			cards = append(cards, value+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
