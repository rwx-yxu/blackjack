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
	DealerPhase()
}

func DealerPhase() {
	//Reveal the dealers hand
	fmt.Printf("Dealer hand: %v%v, %v%v\n", dealer.hand[0].Name, dealer.hand[0].Suit, dealer.hand[1].Name, dealer.hand[1].Suit)

	for {
		if player.score > dealer.score {
			card := deck.Draw()
			dealer.hand = append(dealer.hand, card)
			prescore := dealer.score + card.Value

			for i := 0; i < len(dealer.hand); i++ {
				if prescore > 21 && dealer.hand[i].Name == Ace && dealer.hand[i].Value == 10 {
					dealer.hand[i].Value = 1
					prescore -= 9
				}
			}

			dealer.score = prescore
			fmt.Printf("Card drawn: %v%v\n", card.Name, card.Suit)
			fmt.Printf("Dealer Score: %v\n", dealer.score)

			if dealer.score > 21 {
				fmt.Println("Dealer bust")
				fmt.Println("Player wins!")
				break
			}
		} else if dealer.score > player.score {
			fmt.Println("Dealer wins!")
			break
		} else {
			fmt.Println("Draw. Game tied.")
			break
		}
	}
	os.Exit(0)
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
			prescore := player.score + card.Value

			for i := 0; i < len(player.hand); i++ {
				if prescore > 21 && player.hand[i].Name == Ace && player.hand[i].Value == 10 {
					player.hand[i].Value = 1
					prescore -= 9
				}
			}

			player.score = prescore

			fmt.Printf("Card drawn: %v%v\n", card.Name, card.Suit)
			fmt.Printf("Score: %v\n", player.score)
			if player.score > 21 {
				GameOver()
			}
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
