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
	m := map[CardName]int{
		Ace:   0,
		Two:   0,
		Three: 0,
		Four:  0,
		Five:  0,
		Six:   0,
		Seven: 0,
		Eight: 0,
		Nine:  0,
		Ten:   0,
		Jack:  0,
		Queen: 0,
		King:  0,
	}

	for len(d.Cards) < 52 {
		//Generate a random index of Card Name to be added to deck
		rand.Seed(time.Now().UnixNano())
		randIndx := rand.Intn(len(CardNames))
		cardName := CardNames[randIndx]

		val, _ := m[cardName]

		if val == 4 {
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
		}

		m[cardName] += 1
		d.AddCard(c)
	}
	return d
}

func (d *Deck) AddCard(c Card) {
	d.Cards = append(d.Cards, c)
}
