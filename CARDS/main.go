package main

//import "fmt"

func main() {
	/* 	cards := newDeck()

	   	hand, remainingCards := deal(cards, 5)

	   	hand.print()
	   	remainingCards.print() */

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

}
