package main

import "fmt"

//import "fmt"

func main() {
	/* 	cards := newDeck()

	   	hand, remainingCards := deal(cards, 5)

	   	hand.print()
	   	remainingCards.print() */

	cards := newDeckFromFile("my_cards")
	fmt.Println(cards)
	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

}
