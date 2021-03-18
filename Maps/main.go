package main

import "log"

// maps are immutable
// maps order are random

type Champion struct {
	Name       string
	DamageType string
	DPS        int
}

func main() {
	// String map
	ezrealsMap := make(map[string]string)

	ezrealsMap["mid"] = "Inting"
	ezrealsMap["bot"] = "adc"
	ezrealsMap["jg"] = "penta kill"

	// Int map
	heimersMap := make(map[int]int)
	heimersMap[123] = 321

	// Strut map
	rengarsMap := make(map[string]Champion)

	rengar := Champion{
		Name:       "Rengar",
		DamageType: "Attack Damage",
		DPS:        9999999999,
	}

	// Maps stays the same no matter where in the program it is accessed, it remains constant
	// unlike passing pointers and unsure
	// var oWo float32
	// oWo = 11.1

	rengarsMap["Rengar"] = rengar

	log.Println(ezrealsMap["mid"])
	log.Println(heimersMap[123])
	log.Println(rengarsMap["Rengar"])
	log.Println(rengarsMap["Rengar"].Name)
	log.Println(rengarsMap["Rengar"].DamageType)
	log.Println(rengarsMap["Rengar"].DPS)
}
