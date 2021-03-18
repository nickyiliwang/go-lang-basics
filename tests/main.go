package main

import (
	"errors"
	"log"
	"math"
)

func main() {

	result, err := divide(10, 10)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(float32(math.Inf(0)))
	log.Println("result is ", result)

}

func divide(x, y float32) (float32, error) {

	result := x / y

	if y == 0 {
		return result, errors.New("cannot divide by 0, y cannot be 0")
	} else if x == 0 {
		return result, errors.New("cannot divide by 0, x cannot be 0")
	}

	return result, nil
}
