package blackjack_test

import (
	"testing"

	"github.com/rwx-yxu/blackjack"
	"github.com/rwx-yxu/blackjack/card"
)

func TestBlackJackScore(t *testing.T) {
	cardVal, _ := card.Value(card.King)
	King := card.C{
		Name:  card.King,
		Value: cardVal,
		Suit:  card.Diamond,
	}

	cardVal, _ = card.Value(card.Ace)

	Ace := card.C{
		Name:  card.Ace,
		Value: cardVal,
		Suit:  card.Diamond,
	}
	cardVal, _ = card.Value(card.Queen)
	Queen := card.C{
		Name:  card.Queen,
		Value: cardVal,
		Suit:  card.Diamond,
	}
	cardVal, _ = card.Value(card.Nine)
	Nine := card.C{
		Name:  card.Nine,
		Value: cardVal,
		Suit:  card.Diamond,
	}

	var tests = []struct {
		name  string
		cards []card.C
		want  int
	}{
		{"Ace and King draw should be 21", []card.C{Ace, King}, 21},
		{"King, Queen and Ace draw should be 21", []card.C{King, Queen, Ace}, 21},
		{"Nine, Ace and Ace draw should be 21", []card.C{Nine, Ace, Ace}, 21},
	}

	for _, tt := range tests {
		player := &blackjack.Player{}
		t.Run(tt.name, func(t *testing.T) {
			for _, card := range tt.cards {
				blackjack.Hit(player, card)
			}
			if player.GetScore() != tt.want {
				t.Errorf("error score: got %v, want %v", player.GetScore(), tt.want)
			}
		})
	}
}
