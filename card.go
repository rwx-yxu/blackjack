package blackjack

import (
	"errors"
	"fmt"
)

type Card struct {
	Value int
	Name  CardName
	Suit  CardSuit
}

type CardName string
type CardSuit string

const (
	Diamond CardSuit = "♦️"
	Club    CardSuit = "♣️"
	Heart   CardSuit = "♥️"
	Spade   CardSuit = "♠️"
)

const (
	Ace   CardName = "Ace"
	Two   CardName = "2"
	Three CardName = "3"
	Four  CardName = "4"
	Five  CardName = "5"
	Six   CardName = "6"
	Seven CardName = "7"
	Eight CardName = "8"
	Nine  CardName = "9"
	Ten   CardName = "10"
	Jack  CardName = "Jack "
	Queen CardName = "Queen "
	King  CardName = "King "
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

var CardSuits = []CardSuit{
	Diamond,
	Spade,
	Club,
	Heart,
}

func CardValue(cn CardName) (val int, err error) {
	switch cn {
	case Ace:
		return 10, nil
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
