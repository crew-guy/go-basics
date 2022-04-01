package main

import (
	"fmt"
)

func main() {
	// newCardDeck := newDeck()
	// for _, card := range newCardDeck {
	// 	fmt.Println(card)
	// }
	// deckOne, deckTwo := deal(newCardDeck, 4)
	// deckOne.print()
	// deckTwo.print()

	// fmt.Println(newCardDeck.toString())
	// newCardDeck.saveToFile("new_card_deck")

	deckFromFile := readFromFile("new_card_deck")
	deckFromFile.shuffle()
	fmt.Println(deckFromFile)
}

func newCard() string {
	return "Five of diamonds"
}
