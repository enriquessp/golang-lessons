package main

import (
	"fmt"
)


type bot interface {
	greetings() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {

	e := englishBot{}
	s := spanishBot{}

	fmt.Println("English bot ")
	printGreetings(e)

	fmt.Println("Spanish bot ")
	printGreetings(s)
}

func printGreetings(b bot) {
	fmt.Println(b.greetings())
}

func (englishBot) greetings() string {
	return "Hello"
}

func (spanishBot) greetings() string {
	return "Hola"
}
