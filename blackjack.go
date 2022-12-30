package blackjack

import (
	"fmt"
	"io"
	"os"

	"github.com/rwx-yxu/term"
)

type Dealer struct {
	hand  []Card
	score int
}

type Player struct {
	hand  []Card
	score int
}

type User interface {
	GetHand() []Card
	GetScore() int
	SetScore(s int)
	AddCard(c Card)
}

func (d *Dealer) GetHand() []Card { return d.hand }
func (d *Dealer) GetScore() int   { return d.score }
func (d *Dealer) SetScore(s int)  { d.score = s }
func (d *Dealer) AddCard(c Card)  { d.hand = append(d.hand, c) }

func (p *Player) GetHand() []Card { return p.hand }
func (p *Player) GetScore() int   { return p.score }
func (p *Player) SetScore(s int)  { p.score = s }
func (p *Player) AddCard(c Card)  { p.hand = append(p.hand, c) }

var dealer = &Dealer{}
var deck = NewDeck()
var player = &Player{}

func Run(r io.Reader) {
	//Draw phase
	DrawPhase(player, deck)
	DrawPhase(dealer, deck)

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
		card := deck.Draw()
		if card.Name == Ace && !hasAce {
			hasAce = true
		}

		if card.Name == King || card.Name == Queen || card.Name == King || card.Name == Ten {
			hasTenValue = true
		}
		u.SetScore(u.GetScore() + card.Value)
		u.AddCard(card)
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
			card := d.Draw()
			Hit(dealer, card)
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
			card := d.Draw()
			Hit(player, card)
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

func Hit(u User, card Card) {
	u.AddCard(card)
	prescore := u.GetScore() + card.Value
	handCards := u.GetHand()
	for i := 0; i < len(handCards); i++ {
		if prescore > 21 && handCards[i].Name == Ace && handCards[i].Value == 10 {
			handCards[i].Value = 1
			prescore -= 9
		}
	}

	u.SetScore(prescore)

	fmt.Printf("Card drawn: %v%v\n", card.Name, card.Suit)
}

func PlayerBust(r io.Reader) {
	fmt.Println("Game Over")
	fmt.Printf("Dealer hand was: %v%v, %v%v\n", dealer.hand[0].Name, dealer.hand[0].Suit, dealer.hand[1].Name, dealer.hand[1].Suit)
	fmt.Printf("Dealer score: %v\n", dealer.GetScore())
	EndPrompt(r)
}
