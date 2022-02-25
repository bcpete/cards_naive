package main

var card int

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
