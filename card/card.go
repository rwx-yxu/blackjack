package card

import (
	"errors"
	"fmt"
)

type C struct {
	Value int
	Name  Name
	Suit  Suit
}

type Name string
type Suit string

const (
	Diamond Suit = "♦️"
	Club    Suit = "♣️"
	Heart   Suit = "♥️"
	Spade   Suit = "♠️"
)

const (
	Ace   Name = "Ace"
	Two   Name = "2"
	Three Name = "3"
	Four  Name = "4"
	Five  Name = "5"
	Six   Name = "6"
	Seven Name = "7"
	Eight Name = "8"
	Nine  Name = "9"
	Ten   Name = "10"
	Jack  Name = "Jack "
	Queen Name = "Queen "
	King  Name = "King "
)

var Names = []Name{
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

var Suits = []Suit{
	Diamond,
	Spade,
	Club,
	Heart,
}

func Value(n Name) (val int, err error) {
	switch n {
	case Ace:
		return 11, nil
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
	errStr := fmt.Sprintf("Unable to evaluate card value from given card type: %v", n)
	return 0, errors.New(errStr)
}
