package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

type ContactInfo struct {
	Email   string
	ZipCode int
}

func main() {

	steve := Person{
		FirstName: "Steve",
		LastName:  "Aioki",
		ContactInfo: ContactInfo{
			Email:   "steve@gmail.com",
			ZipCode: 19996,
		},
	}

	lulu := Person{
		FirstName: "Lulu",
		LastName:  "Aioki",
		ContactInfo: ContactInfo{
			Email:   "lulu@gmail.com",
			ZipCode: 202020,
		},
	}

	// ***This is an interesting thing in go,
	// Expected output: FirstName === "Joey"
	// But got: "Steve"
	// Before I knew what pointers are doing/ why we need pointers
	// Maybe pointers are used to point to the value's memory address
	// in order for the updateName function to produce expected output
	steve.updateName("Joey")
	steve.print()
	lulu.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

// Why are we using pointer here:
// Simple: the FirstName value is saved onto another address space than the Person struct
// p in updateName is a copy
// So when we update the FirstName's value we want to update the Person struct address space with a pointer

func (p *Person) updateName(newFirstName string) {
	p.FirstName = newFirstName
}
