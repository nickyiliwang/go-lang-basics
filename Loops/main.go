package main

import "log"

// go only has one loop, the for loop

func main() {
	// always in order
	champions := []string{"Rengar", "Nunu", "Samira", "Kayle", "Annie", "Renenkton"}
	// randomized
	colorMap := make(map[string]string)
	// struct, no colons
	type Champion struct {
		Name       string
		DamageType string
		DPS        int
	}

	colorMap["Red"] = "Red"
	colorMap["Blue"] = "Blue"
	colorMap["Yellow"] = "Yellow"
	colorMap["White"] = "White"
	colorMap["Black"] = "Black"
	colorMap["Orange"] = "Orange"
	colorMap["Rainbow"] = "Rainbow"

	// basic for loop
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// using the index
	for i := 0; i < len(champions); i++ {
		log.Println(champions[i])
	}

	// not using the index
	// similar to forEach
	for i, y := range champions {

		log.Println("champ index", i)
		log.Println("champs", y)
	}

	for i, y := range colorMap {
		log.Println("color index", i)
		log.Println("color", y)
	}

	var champSlice []Champion

	// var coolChamp Champion
	// coolChamp.Name = "Rengar"
	// coolChamp.DamageType = "Attack Damage"
	// coolChamp.DPS = 9000000000

	coolChamp := Champion{
		Name:       "Rengar",
		DamageType: "AD",
		DPS:        90000000000000000,
	}

	okayChamp := Champion{
		Name:       "Akali",
		DamageType: "Mixed",
		DPS:        90000,
	}

	champSlice = append(champSlice, coolChamp)
	champSlice = append(champSlice, okayChamp)

	log.Println(champSlice)

	// for champ := 0; champ < len(champSlice); champ++ {
	// 	log.Println("Champion", champSlice[champ])
	// }

	for i, x := range champSlice {
		log.Println(i, x.Name, x.DPS)
	}

}
