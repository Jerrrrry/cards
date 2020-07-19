package main

func main() {
	//var card string = "Ace of Spades"
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// //cards.print()
	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
