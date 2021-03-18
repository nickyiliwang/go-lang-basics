package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// fmt.Println("hello world")

	var say1 string = saySomething("Too deep for me")
	var say2 string = saySomething("WOWZERS deep for me")

	fmt.Println(say1)
	fmt.Println(say2)

}

func saySomething(s string) string {
	var stuff = s + " " + quote.Go()
	return stuff
}

