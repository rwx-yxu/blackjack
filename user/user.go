package user

import (
	"fmt"

	"github.com/rwx-yxu/blackjack/card"
)

type Dealer struct {
	*U
}

type Player struct {
	*U
}

type U struct {
	Hand  []card.C
	Score int
}

func (d Dealer) ShowPartialHand() {
	fmt.Printf("Dealer cards: ?, %v%v\n", d.Hand[1].Name, d.Hand[1].Suit)
}

func (user U) ShowHand() string {
	return fmt.Sprintf("%v%v, %v%v", user.Hand[0].Name, user.Hand[0].Suit, user.Hand[1].Name, user.Hand[1].Suit)
}

func (user *U) AddCard(c card.C) { user.Hand = append(user.Hand, c) }

func (user *U) Hit(c card.C) {
	user.AddCard(c)
	prescore := user.Score + c.Value
	handCards := user.Hand
	for i := 0; i < len(handCards); i++ {
		if prescore > 21 && handCards[i].Name == card.Ace && handCards[i].Value == 11 {
			handCards[i].Value = 1
			prescore -= 10
		}
	}

	user.Score = prescore
}
