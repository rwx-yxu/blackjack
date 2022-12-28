package blackjack

import (
	"fmt"
	"io"
)

func Run(in io.Reader) {
	deck := NewDeck()
	fmt.Printf("Deck generated. Deck size: %v\n", len(deck.Cards))
}
