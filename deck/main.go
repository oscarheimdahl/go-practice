package main

func main() {
	cards := newDeck()
	cards.print()
	cardsS := shuffle(cards)
	cardsS.print()
}
