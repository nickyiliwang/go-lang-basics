package main

import (
	"fmt"
	"time"
)

// Interface can be used to reduce redundant use of identical functions

type bot interface {
	// FN         return type
	getGreeting() string
}
type englishBot struct {
	country string
	date    time.Time
}
type spanishBot struct{}
type badBot struct{}

func main() {
	// declaring an empty struct
	eb := englishBot{
		country: "USA",
		date:    time.Now(),
	}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// bb := badBot{}
	// cannot use bb (variable of type badBot) as bot value in argument to printGreeting: missing method getGreetingcompilerInvalidIfaceAssign
	// printGreeting(bb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// custom logic for generating an english getGreeting
	fmt.Println(eb.country)
	fmt.Println(eb.date)
	return "Hello there!"
}

// can emit the receiver value if we aren't using it
func (spanishBot) getGreeting() string {
	return "Hola Amigo!"
}
