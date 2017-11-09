package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new typ of 'deck'
// wich is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//d copy of deck working with can be called by d, declaring deck means every variable of type deck can call the function on itself cards. print is the name of the method.
// ex: cards.print() calls the print method which executes the for loop, d is comparable to this or self in c#
//d declaration by convention is a 1-3 letter abbreviation that matches type in go, not required though
//receiver comparable to function in a class
func (d deck) print() { //receiver
	for i, card := range d {
		fmt.Println(i, card)
	}

}

//Create a hand and leftover cards both of type deck, returning two values
//using range of slice going form start to handsize and remaining from excluded to end of slice as range 2nd value is not inclusive (shorthand) [0,variable]==[:variable] and [handsize:lastelementinclusive]==[variable:]
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Implement receiver to call as method in this case
//utilize strings package to use join function to create one string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//converted slice of strings to one string with toString function, then convert to byte slice in write file method
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
