package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Assuming you know ahead of time the format you are going to receive, so we can build a strcut around it

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Breed     string `json:"breed"`
	IsChamp   bool   `json:"isChamp"`
}

func main() {
	jsonAPI := `[
		{
		   "first_name":"nick",
		   "last_name":"wang",
		   "breed":"human",
		   "isChamp":false
		},
		{
		   "first_name":"rengar",
		   "last_name":"mclovin",
		   "breed":"dog",
		   "isChamp":true
		}
	 ]`

	//  golang functions can be marshalled or unmarshalled, json we get from api is unmarshalled
	// unmarshalled := []Person()
	var unmarshalled []Person

	// JSON => Struct
	err := json.Unmarshal([]byte(jsonAPI), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var personSlice []Person

	var p1 Person
	p1.FirstName = "Iron"
	p1.LastName = "Man"
	p1.Breed = "Human"
	p1.IsChamp = false

	p2 := Person{
		FirstName: "Dr.",
		LastName:  "Mundo",
		Breed:     "Freak",
		IsChamp:   true,
	}

	personSlice = append(personSlice, p1)
	personSlice = append(personSlice, p2)

	log.Println(personSlice)

	marshalToJSON, err := json.MarshalIndent(personSlice, "", "	")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(marshalToJSON))
}
