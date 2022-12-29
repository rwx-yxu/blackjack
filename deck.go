package blackjack

import (
	"log"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	d := new(Deck)

	deckTracker := map[CardName]map[CardSuit]bool{
		Ace: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Two: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Three: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Four: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Five: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Six: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Seven: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Eight: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Nine: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
		Ten: {
			Diamond: false,
			Heart:   false,
			Club:    false,
			Spade:   false,
		},
		Jack: {
			Diamond: false,
			Spade:   false,
			Club:    false,
			Heart:   false,
		},
		Queen: {
			Diamond: false,
			Club:    false,
			Spade:   false,
			Heart:   false,
		},
		King: {
			Diamond: false,
			Club:    false,
			Heart:   false,
			Spade:   false,
		},
	}

	for len(d.Cards) < 52 {
		//Generate a random index of card name and card suit to be added to deck
		rand.Seed(time.Now().UnixNano())
		randIndx := rand.Intn(len(CardNames))
		cardName := CardNames[randIndx]

		randIndx = rand.Intn(len(CardSuits))
		cardSuit := CardSuits[randIndx]

		val, _ := deckTracker[cardName][cardSuit]

		if val {
			continue
		}

		//Initialize card and add to deck
		cardVal, err := CardValue(cardName)
		if err != nil {
			//Passed invalid
			log.Print(err)
			//Remove invalid card name from CardNames array and continue
			CardNames = append(CardNames[:randIndx], CardNames[randIndx+1:]...)
			continue
		}
		c := Card{
			Value: cardVal,
			Name:  cardName,
			Suit:  cardSuit,
		}

		deckTracker[cardName][cardSuit] = true
		d.AddCard(c)
	}
	return d
}

func (d *Deck) AddCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Draw() Card {
	//Check that deck size is empty. If empty, generate a new deck.
	if len(d.Cards) == 0 {
		d = NewDeck()
	}

	topCard := d.Cards[len(d.Cards)-1]
	//Remove last element from cards list
	d.Cards = d.Cards[:len(d.Cards)-1]
	return topCard
}
