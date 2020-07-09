package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	cards.saveToFile("cardfile.txt")

	moreCards := deckFromFile("cardfile.txt")
	moreCards.print()
}
