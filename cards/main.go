package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// cards := newDeck()

	//cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("---------")
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
