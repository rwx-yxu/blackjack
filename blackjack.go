package blackjack

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Value int
	Name  CardName
}

type CardName string

const (
	Ace   CardName = "Ace"
	Two   CardName = "Two"
	Three CardName = "Three"
	Four  CardName = "Four"
	Five  CardName = "Five"
	Six   CardName = "Six"
	Seven CardName = "Seven"
	Eight CardName = "Eight"
	Nine  CardName = "Nine"
	Ten   CardName = "Ten"
	Jack  CardName = "Jack"
	Queen CardName = "Queen"
	King  CardName = "King"
)

var CardNames = []CardName{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
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

func CardValue(cn CardName) (val int, err error) {
	switch cn {
	case Ace:
		return 1, nil
	case Two:
		return 2, nil
	case Three:
		return 3, nil
	case Four:
		return 4, nil
	case Five:
		return 5, nil
	case Six:
		return 6, nil
	case Seven:
		return 7, nil
	case Eight:
		return 8, nil
	case Nine:
		return 9, nil
	case Ten:
		return 10, nil
	case Jack:
		return 10, nil
	case Queen:
		return 10, nil
	case King:
		return 10, nil
	}
	errStr := fmt.Sprintf("Unable to evaluate card value from given card type: %v", cn)
	return 0, errors.New(errStr)
}

func (d *Deck) AddCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func Run(in io.Reader) {
	deck := NewDeck()
	fmt.Printf("Deck generated. Deck size: %v\n", len(deck.Cards))
}
