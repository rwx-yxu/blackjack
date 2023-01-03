package blackjack_test

import (
	"testing"

	"github.com/rwx-yxu/blackjack"
	"github.com/rwx-yxu/blackjack/user"
)

func TestPlayerStartingHandCount(t *testing.T) {
	deck := blackjack.NewDeck()
	var player = user.Player{
		U: &user.U{},
	}

	blackjack.DrawPhase(player.U, deck)

	if len(player.Hand) != 2 {
		t.Errorf("starting hand should be 2: got %v, want %v", len(player.Hand), 2)
	}

}
