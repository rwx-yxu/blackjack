package blackjack

import (
	"fmt"
	"io"
)

type Dealer struct {
	hand  []Card
	score int
}

type Player struct {
	hand  []Card
	score int
}

func Run(in io.Reader) {
	dealer := Dealer{}
	deck := NewDeck()
	player := Player{}

	//Draw phase
	//Deal 2 cards each so 2 loops
	for i := 0; i < 2; i++ {
		player.hand = append(player.hand, deck.Draw())
		dealer.hand = append(dealer.hand, deck.Draw())
	}

	fmt.Printf("Dealer cards: ?, %v%v\n", dealer.hand[1].Name, dealer.hand[1].Suit)

	fmt.Printf("Player cards: %v%v, %v%v\n", player.hand[0].Name, player.hand[0].Suit, player.hand[1].Name, player.hand[1].Suit)
	//Player phase
}

/*
func DrawPhase(p Player,d Dealer, deck Deck) ([]Cards, []Cards){
  //Deal 2 cards each so 2 loops
  for i := 0; i < 2; i++ {
    p.hand = append(p.hand, deck.Draw())
    d.hand = append(d.hand, deck.Draw())
  }
	return p.hand,d.hand
}

*/
