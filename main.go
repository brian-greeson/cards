package main

func main() {
	cards := newDeck()

	cards.print()

	// hand, remainingCards := deal(cards, 5)

	cards.saveToFile("cardfile.txt")

}
