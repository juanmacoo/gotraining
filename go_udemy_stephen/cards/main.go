package main

func main() {
	cards := newDeck()
	cards = shuffle(cards)
	cards.print()

	// cards.saveToFile()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}
