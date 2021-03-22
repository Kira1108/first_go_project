package main

import "fmt"

func main() {
	cards := newDeckFromFile("cards.txt")
	fmt.Println("========Before shuffle======")
	cards.print()
	cards.shuffle()
	fmt.Println("========After shuffle =======")
	cards.print()
	// cards := newDeck()
	// cards.saveToFile("cards.txt")
	// fmt.Println(cards.toString())
	// hand, remain := deal(cards, 5)
	// fmt.Println("---------Hand----------")
	// hand.print()
	// fmt.Println("---------Remain----------")
	// remain.print()

}
