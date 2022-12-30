package blackjack_test

import (
	"testing"

	"github.com/rwx-yxu/blackjack"
)

func TestDeckCardCount(t *testing.T) {
	deck := blackjack.NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("Deck card count incorrect, got %v, want: %v", len(deck.Cards), 52)
	}
}
