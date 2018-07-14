package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// cards := [] string {"Ace of Spades", newCard()}
	// cards := deck {"Ace of Spades", newCard()}
	cards := newDeck()

	// Adding value to an array. Actually creating a new arry appending a new vaue
	// cards = append(cards, "Six of Spades")

	// Iterating
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//	cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Hand: ")
	// hand.print()

	// fmt.Println("Ramaining cards")
	// remainingDeck.print()

	//  reassigning
	//	card = "Five of Diamonds"

	// fmt.Println(card)
	// fmt.Println(cards)
	// printState()

	fmt.Println(cards.toString())

	fmt.Println("SavingFile")
	cards.saveToFile()

	fmt.Println("Reading a file")
	cardsFromFile := newDeckFromFile("file_writen_by_go_program")
	cardsFromFile.print()

	fmt.Println("Shuffling")
	cardsFromFile.shuffle()
	cardsFromFile.print()
}

func newCard() string {
	return "Five of Diamonds"
}
