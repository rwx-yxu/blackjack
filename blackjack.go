package blackjack

import (
	"fmt"
	"io"
	"os"

	"github.com/rwx-yxu/blackjack/user"
	"github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
)

var dealer = user.Dealer{
	U: &user.U{},
}
var deck = NewDeck()
var player = user.Player{
	U: &user.U{},
}

func Run(r io.Reader) {
	sequence.OnIfTerminal(r)
	fmt.Println(sequence.CLSEntire)
	//Draw phase
	DrawPhase(dealer.U, deck)
	DrawPhase(player.U, deck)
	fmt.Printf("Deck size remaining: %v\n", len(deck.Cards))
	dealer.ShowPartialHand()
	fmt.Printf("Player cards: %v\n", player.ShowHand())

	fmt.Printf("Score: %v\n", player.Score)
	fmt.Println("----------------------------------")
	if player.Score != 21 {
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
			dealer = user.Dealer{
				U: &user.U{},
			}
			player = user.Player{
				U: &user.U{},
			}
			Run(r)
		case "2":
			fmt.Println("Thank you for playing.")
			os.Exit(0)
		default:
			fmt.Println("Please enter a valid response of '1' or '2'")
		}
	}
}

func DrawPhase(u *user.U, deck *Deck) {

	for i := 0; i < 2; i++ {
		if len(deck.Cards) == 0 {
			deck = NewDeck()
		}
		c := deck.Draw()
		u.Hit(c)
	}

}

func DealerPhase() {
	//Reveal the dealers hand
	fmt.Printf("Dealer hand: %v\n", dealer.ShowHand())

	for {
		if player.Score > dealer.Score {
			if len(deck.Cards) == 0 {
				deck = NewDeck()
			}
			c := deck.Draw()
			dealer.Hit(c)
			fmt.Printf("Card drawn: %v%v. Deck size remaining:%v\n", c.Name, c.Suit, len(deck.Cards))
			fmt.Printf("Dealer score: %v\n", dealer.Score)
			if dealer.Score > 21 {
				fmt.Println("Dealer bust")
				fmt.Println("Player wins!")
				break
			}
		} else if dealer.Score > player.Score {
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
			fmt.Printf("Standing with a total of: %v\n", player.Score)
			return
		case "2":
			if len(deck.Cards) == 0 {
				deck = NewDeck()
			}
			c := deck.Draw()
			player.Hit(c)
			fmt.Printf("Card drawn: %v%v. Deck size remaining:%v\n", c.Name, c.Suit, len(deck.Cards))
			fmt.Printf("Player score: %v\n", player.Score)
			if player.Score > 21 {
				PlayerBust(r)
			}
			break
		default:
			fmt.Println("Please enter a valid response of '1' or '2'")
		}
	}
}

func PlayerBust(r io.Reader) {
	fmt.Println("----------------------------------")
	fmt.Println("Game Over")
	fmt.Printf("Dealer hand was: %v\n", dealer.ShowHand())
	fmt.Printf("Dealer score: %v\n", dealer.Score)
	EndPrompt(r)
}
