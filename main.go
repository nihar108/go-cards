package main

func main() {
	cards := newDeckFromFile("cardsfile")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("cardsfile")
	cards.shuffle()
	cards.print()
}
