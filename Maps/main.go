package main

import (
	"fmt"
)

func main() {

	// Init Option 1
	// value := map[key type]values type
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf213",
		"black": "#ffffff",
	}

	// Init Option 2
	colors := make(map[string]string)
	// maps always uses squares syntax
	colors["white"] = "ffffff"
	// deleting
	delete(colors, "white")
	colors["red"] = "#fof0f0f"

	fmt.Println(colors1)
	fmt.Println(colors)

	printMap(colors1)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}
