package user

import (
	"fmt"

	"github.com/rwx-yxu/blackjack/card"
)

type Dealer struct {
	hand  []card.C
	score int
}

type Player struct {
	hand  []card.C
	score int
}

type U interface {
	Hand() []card.C
	Score() int
	SetScore(s int)
	AddCard(c card.C)
}

func (d *Dealer) Hand() []card.C   { return d.hand }
func (d *Dealer) Score() int       { return d.score }
func (d *Dealer) SetScore(s int)   { d.score = s }
func (d *Dealer) AddCard(c card.C) { d.hand = append(d.hand, c) }
func (d *Dealer) ShowPartialHand() {
	fmt.Printf("Dealer cards: ?, %v%v\n", d.hand[1].Name, d.hand[1].Suit)
}

func (p *Player) Hand() []card.C   { return p.hand }
func (p *Player) Score() int       { return p.score }
func (p *Player) SetScore(s int)   { p.score = s }
func (p *Player) AddCard(c card.C) { p.hand = append(p.hand, c) }

func ShowHand(u U) string {
	return fmt.Sprintf("%v%v, %v%v", u.Hand()[0].Name, u.Hand()[0].Suit, u.Hand()[1].Name, u.Hand()[1].Suit)
}
