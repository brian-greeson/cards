package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	faces := []string{"Ace", "King", "Queen", "Jack"}

	for _, suit := range suits {
		for _, face := range faces {
			cards = append(cards, face+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
