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

var dealer = Dealer{}
var deck = NewDeck()
var player = Player{}

func Run(r io.Reader) {
	//Draw phase
	//Deal 2 cards each so 2 loops
	for i := 0; i < 2; i++ {
		card := deck.Draw()
		player.score += card.Value
		player.hand = append(player.hand, card)
		card = deck.Draw()
		dealer.score += card.Value
		dealer.hand = append(dealer.hand, card)
	}

	fmt.Printf("Dealer cards: ?, %v%v\n", dealer.hand[1].Name, dealer.hand[1].Suit)

	fmt.Printf("Player cards: %v%v, %v%v\n", player.hand[0].Name, player.hand[0].Suit, player.hand[1].Name, player.hand[1].Suit)
	fmt.Printf("Score: %v\n", player.score)
	PlayerPhase(r)
}

func PlayerPhase(r io.Reader) {
	for {
		//Player phase
		fmt.Println("(1) Stand or (2) Draw?")
		resp, _ := term.Prompt(r, ">")
		switch resp {
		case "1":
			fmt.Printf("Standing with a total of: %v\n", player.score)
			return
		case "2":
			card := deck.Draw()
			player.hand = append(player.hand, card)

			if card.Name == Ace && (player.score+card.Value > 21) {
				card.Value = 1
			}

			player.score += card.Value

			fmt.Printf("Card drawn: %v%v\n", card.Name, card.Suit)
			if player.score > 21 {
				GameOver()
			}
			fmt.Printf("Score: %v\n", player.score)
		default:
			fmt.Println("Please enter a valid response of '1' or '2'")
		}
	}
}

func GameOver() {
	fmt.Println("Game Over")
	fmt.Printf("Dealer hand was: %v%v, %v%v\n", dealer.hand[0].Name, dealer.hand[0].Suit, dealer.hand[1].Name, dealer.hand[1].Suit)
	os.Exit(0)
}
