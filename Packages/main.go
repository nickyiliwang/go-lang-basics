package main

import (
	"log"

	"github.com/nickyiliwang/champion"
)

func main() {
	rengar := champion.Champion{
		Name:       "Rengar",
		Role:       "Assassin",
		DamageType: "AD",
		DPS:        90000000000000,
	}

	log.Println("Hello", rengar)

}
