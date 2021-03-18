package main

import (
	"log"

	"github.com/nickyiliwang/random"
)

// Communication between programs:
// 1. Creating an importing a funcion
// 2. Using a Pointer and receiver function
// 3. Using a Channel

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := random.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// channel that only accepts integers
	intChan := make(chan int)

	// close the channel as soon as the function executes, CalculateValue in this case
	defer close(intChan)

	// Passing value into the channel
	go CalculateValue(intChan)

	// getting value from the channel and assigning into num
	num := <-intChan

	log.Print("number from intChan ", num)

}
