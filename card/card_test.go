package card_test

import (
	"fmt"

	"github.com/rwx-yxu/blackjack/card"
)

func ExampleCardNames() {
	for _, n := range card.Names {
		fmt.Printf("%q\n", n)
	}

	//Output:
	//"Ace"
	//"2"
	//"3"
	//"4"
	//"5"
	//"6"
	//"7"
	//"8"
	//"9"
	//"10"
	//"Jack "
	//"Queen "
	//"King "

}

func ExampleCardValues() {
	for _, n := range card.Names {
		v, _ := card.Value(n)
		fmt.Printf("%v\n", v)
	}

	//Output:
	//11
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10
	//10
	//10
	//10

}
