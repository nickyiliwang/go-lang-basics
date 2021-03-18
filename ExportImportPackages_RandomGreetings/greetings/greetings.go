// functional module exported as package

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
// func name    param type & return type
func Hello(name string) (string, error) {
	// If no name was given, return an error with message
	if name == "" {
		return "", errors.New("empty name, empty life")
	}

	// Return a greeting that embeds the name in a message
	// := is a variable declaration
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())

}

// randomFormat returns one of a set of greeting messages. The return
// message is selected at random
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to meet you, %v!",
		"All Hail the mighty %v!",
	}

	// return a randomly selected message format by specifying a
	// random index for the slice of formats
	// len(formats) is the length of formats
	return formats[rand.Intn(len(formats))]
}
