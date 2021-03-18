package main

import "log"

func main() {
	var isFunny bool = false

	if isFunny {
		log.Println("I am Reallly funny like really funny!!")
	} else {
		log.Println("WeirdChamp")
	}

	switch isFunny {
	case true:
		log.Println("heheheheheehh")
	case false:
		log.Println("sadge")
	default:
		log.Println("default")
}

}
