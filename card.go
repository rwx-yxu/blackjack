package blackjack

import (
	"errors"
	"fmt"
)

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
