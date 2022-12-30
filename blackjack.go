package blackjack

import (
	"fmt"
	"io"
	"os"

	"github.com/rwx-yxu/blackjack/card"
	"github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
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

func (p *Player) GetHand() []card.C { return p.hand }
func (p *Player) GetScore() int     { return p.score }
func (p *Player) SetScore(s int)    { p.score = s }
func (p *Player) AddCard(c card.C)  { p.hand = append(p.hand, c) }

var dealer = &Dealer{}
var deck = NewDeck()
var player = &Player{}

func Run(r io.Reader) {
	sequence.OnIfTerminal(r)
	fmt.Println(sequence.CLSEntire)
	//Draw phase
	DrawPhase(player, deck)
	DrawPhase(dealer, deck)
	fmt.Printf("Deck size remaining: %v\n", len(deck.Cards))
	fmt.Printf("Dealer cards: ?, %v%v\n", dealer.hand[1].Name, dealer.hand[1].Suit)

	fmt.Printf("Player cards: %v%v, %v%v\n", player.hand[0].Name, player.hand[0].Suit, player.hand[1].Name, player.hand[1].Suit)

	fmt.Printf("Score: %v\n", player.score)

	PlayerPhase(r)
	DealerPhase()
	EndPrompt(r)

}

func EndPrompt(r io.Reader) {
	for {
		fmt.Println("(1) Play again or (2) End game?")
		resp, _ := term.Prompt(r, ">")
		switch resp {
		case "1":
			//Reset hand and scores. Deck should remain the same
			dealer = &Dealer{}
			player = &Player{}
			Run(r)
		case "2":
			fmt.Println("Thank you for playing.")
			os.Exit(0)
		default:
			fmt.Println("Please enter a valid response of '1' or '2'")
		}
	}
}
func DrawPhase(u User, deck *Deck) {

	hasAce := false
	hasTenValue := false

	for i := 0; i < 2; i++ {
		if len(deck.Cards) == 0 {
			deck = NewDeck()
		}
		c := deck.Draw()
		if c.Name == card.Ace && !hasAce {
			hasAce = true
		}

		if c.Name == card.King || c.Name == card.Queen || c.Name == card.King || c.Name == card.Ten {
			hasTenValue = true
		}
		u.SetScore(u.GetScore() + c.Value)
		u.AddCard(c)
	}

	if hasAce && hasTenValue {
		u.SetScore(21)
	}
}

func DealerPhase() {
	//Reveal the dealers hand
	fmt.Printf("Dealer hand: %v%v, %v%v\n", dealer.hand[0].Name, dealer.hand[0].Suit, dealer.hand[1].Name, dealer.hand[1].Suit)

	for {
		if player.GetScore() > dealer.GetScore() {
			if len(deck.Cards) == 0 {
				deck = NewDeck()
			}
			c := deck.Draw()
			Hit(dealer, c)
			fmt.Printf("Dealer score: %v\n", dealer.GetScore())
			if dealer.score > 21 {
				fmt.Println("Dealer bust")
				fmt.Println("Player wins!")
				break
			}
		} else if dealer.GetScore() > player.GetScore() {
			fmt.Println("Dealer wins!")
			break
		} else {
			fmt.Println("Draw. Game tied.")
			break
		}
	}
}

func PlayerPhase(r io.Reader) {
	for {
		//Player phase
		fmt.Println("(1) Stand or (2) Hit?")
		resp, _ := term.Prompt(r, ">")
		switch resp {
		case "1":
			fmt.Printf("Standing with a total of: %v\n", player.GetScore())
			return
		case "2":
			if len(deck.Cards) == 0 {
				deck = NewDeck()
			}
			c := deck.Draw()
			Hit(player, c)
			fmt.Printf("Player score: %v\n", player.GetScore())
			if player.GetScore() > 21 {
				PlayerBust(r)
			}
			break
		default:
			fmt.Println("Please enter a valid response of '1' or '2'")
		}
	}
}

func Hit(u User, c card.C) {
	u.AddCard(c)
	prescore := u.GetScore() + c.Value
	handCards := u.GetHand()
	for i := 0; i < len(handCards); i++ {
		if prescore > 21 && handCards[i].Name == card.Ace && handCards[i].Value == 10 {
			handCards[i].Value = 1
			prescore -= 9
		}
	}

	u.SetScore(prescore)

	fmt.Printf("Card drawn: %v%v. Deck size remaining:%v\n", c.Name, c.Suit, len(deck.Cards))
}

func PlayerBust(r io.Reader) {
	fmt.Println("Game Over")
	fmt.Printf("Dealer hand was: %v%v, %v%v\n", dealer.hand[0].Name, dealer.hand[0].Suit, dealer.hand[1].Name, dealer.hand[1].Suit)
	fmt.Printf("Dealer score: %v\n", dealer.GetScore())
	EndPrompt(r)
}
