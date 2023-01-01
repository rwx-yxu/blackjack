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

type User interface {
	GetHand() []card.C
	GetScore() int
	SetScore(s int)
	AddCard(c card.C)
}

func (d *Dealer) GetHand() []card.C { return d.hand }
func (d *Dealer) GetScore() int     { return d.score }
func (d *Dealer) SetScore(s int)    { d.score = s }
func (d *Dealer) AddCard(c card.C)  { d.hand = append(d.hand, c) }
func (d *Dealer) ShowPartialHand() {
	fmt.Printf("Dealer cards: ?, %v%v\n", dealer.hand[1].Name, dealer.hand[1].Suit)
}

func (p *Player) GetHand() []card.C { return p.hand }
func (p *Player) GetScore() int     { return p.score }
func (p *Player) SetScore(s int)    { p.score = s }
func (p *Player) AddCard(c card.C)  { p.hand = append(p.hand, c) }

func ShowHand(u User) String {
	return fmt.Sprintf("%v%v, %v%v", u.GetHand()[0].Name, u.GetHand()[0].Suit, u.GetHand()[1].Name, u.GetHand()[1].Suit)
}
