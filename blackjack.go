package blackjack

import (
	"fmt"
	"io"
	"os"

	"github.com/rwx-yxu/blackjack/card"
	"github.com/rwx-yxu/blackjack/user"
	"github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
)

var dealer = &user.Dealer{}
var deck = NewDeck()
var player = &user.Player{}

func Run(r io.Reader) {
	sequence.OnIfTerminal(r)
	fmt.Println(sequence.CLSEntire)
	//Draw phase
	DrawPhase(player, deck)
	DrawPhase(dealer, deck)
	fmt.Printf("Deck size remaining: %v\n", len(deck.Cards))
	dealer.ShowPartialHand()

	fmt.Printf("Player cards: %v\n", user.ShowHand(player))

	fmt.Printf("Score: %v\n", player.GetScore())
	fmt.Println("----------------------------------")
	if player.score != 21 {
		PlayerPhase(r)
	} else {
		fmt.Println("Blackjack Hit! Standing with current cards.")
	}
	fmt.Println("----------------------------------")
	DealerPhase()
	fmt.Println("----------------------------------")
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

	for i := 0; i < 2; i++ {
		if len(deck.Cards) == 0 {
			deck = NewDeck()
		}
		c := deck.Draw()
		u.SetScore(u.GetScore() + c.Value)
		u.AddCard(c)
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
			fmt.Printf("Card drawn: %v%v. Deck size remaining:%v\n", c.Name, c.Suit, len(deck.Cards))
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
			fmt.Printf("Card drawn: %v%v. Deck size remaining:%v\n", c.Name, c.Suit, len(deck.Cards))
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
		if prescore > 21 && handCards[i].Name == card.Ace && handCards[i].Value == 11 {
			handCards[i].Value = 1
			prescore -= 10
		}
	}

	u.SetScore(prescore)

}

func PlayerBust(r io.Reader) {
	fmt.Println("----------------------------------")
	fmt.Println("Game Over")
	fmt.Printf("Dealer hand was: %v\n", user.ShowHand(dealer))
	fmt.Printf("Dealer score: %v\n", dealer.GetScore())
	EndPrompt(r)
}
