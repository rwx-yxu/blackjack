package blackjack_test

import (
	"testing"

	"github.com/rwx-yxu/blackjack"
	"github.com/rwx-yxu/blackjack/card"
)

func TestDeckCardCount(t *testing.T) {
	deck := blackjack.NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("Deck card count incorrect, got %v, want: %v", len(deck.Cards), 52)
	}
}

func TestDeckCardCountAfterDraw(t *testing.T) {
	deck := blackjack.NewDeck()
	deck.Draw()

	if len(deck.Cards) != 51 {
		t.Errorf("Deck card count after draw incorrect, got %v, want: %v", len(deck.Cards), 51)
	}

}

func TestDeckDrawCard(t *testing.T) {
	deck := blackjack.NewDeck()
	card := deck.Draw()
	testStruct(card, t)

}

func testStruct(x any, t *testing.T) {
	switch x.(type) {
	case card.C:
		break
	default:
		t.Errorf("Card drawn is not of type card.C")
	}
}
