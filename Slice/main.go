package main

import (
	"log"
	"sort"
)

func main() {
	// variable
	// var fish string = "fish"

	// variable slice
	var dogs []string

	dogs = append(dogs, "Golden Retriver")
	dogs = append(dogs, "Chu Wa Wa")
	dogs = append(dogs, "Mutt")
	dogs = append(dogs, "Hatchi")

	// sortedDoggos :=
	// sort.Ints(dogs)

	// slices shorthand
	champions := []string{"Rengar", "Soraka", "Akali", "Nunu", "Vayne", "Kha'Zix"}

	sort.Strings(dogs)
	log.Println("Doggos:", dogs)
	log.Println("the 3rd doggo is:", dogs[2])
	log.Println("Champs I play the most:", champions)
	// only display the first 3 strings [0 <= start of slice, 3 <= end of slice]
	log.Println("Top 3 Champs I play:", champions[0:3])
}
