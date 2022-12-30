package blackjack

import (
	"log"
	"math/rand"
	"time"

	"github.com/rwx-yxu/blackjack/card"
)

type Deck struct {
	Cards []card.C
}

func NewDeck() *Deck {
	d := new(Deck)

	deckTracker := map[card.Name]map[card.Suit]bool{
		card.Ace: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Two: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Three: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Four: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Five: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Six: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Seven: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Eight: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Nine: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
		card.Ten: {
			card.Diamond: false,
			card.Heart:   false,
			card.Club:    false,
			card.Spade:   false,
		},
		card.Jack: {
			card.Diamond: false,
			card.Spade:   false,
			card.Club:    false,
			card.Heart:   false,
		},
		card.Queen: {
			card.Diamond: false,
			card.Club:    false,
			card.Spade:   false,
			card.Heart:   false,
		},
		card.King: {
			card.Diamond: false,
			card.Club:    false,
			card.Heart:   false,
			card.Spade:   false,
		},
	}

	for len(d.Cards) < 52 {
		//Generate a random index of card name and card suit to be added to deck
		rand.Seed(time.Now().UnixNano())
		randIndx := rand.Intn(len(card.Names))
		cardName := card.Names[randIndx]

		randIndx = rand.Intn(len(card.Suits))
		cardSuit := card.Suits[randIndx]

		val, _ := deckTracker[cardName][cardSuit]

		if val {
			continue
		}

		//Initialize card and add to deck
		cardVal, err := card.Value(cardName)
		if err != nil {
			//Passed invalid
			log.Print(err)
			//Remove invalid card name from CardNames array and continue
			card.Names = append(card.Names[:randIndx], card.Names[randIndx+1:]...)
			continue
		}
		c := card.C{
			Value: cardVal,
			Name:  cardName,
			Suit:  cardSuit,
		}

		deckTracker[cardName][cardSuit] = true
		d.AddCard(c)
	}
	return d
}

func (d *Deck) AddCard(c card.C) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Draw() card.C {
	topCard := d.Cards[len(d.Cards)-1]
	//Remove last element from cards list
	d.Cards = d.Cards[:len(d.Cards)-1]
	return topCard
}
