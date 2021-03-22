package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func newDeck() deck {
	// define slices
	cardSuits := []string{"Spade", "Heart", "Diamond", "Club"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight"}

	// initialize a deck
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append values
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filepath string) error {
	return ioutil.WriteFile(filepath, []byte(d.toString()), 0666)
}

func newDeckFromFile(filepath string) deck {
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error: An error occured while reading the file.")
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
