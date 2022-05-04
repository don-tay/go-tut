package main

func main() {
	deck := readDeck("my_cards.txt")
	deck.shuffle()
	deck.printDeck()
	deck.saveDeck("my_cards.txt")
}
