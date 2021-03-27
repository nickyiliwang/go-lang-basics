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

	// steve.updateName("Joey")
	// steve.print()
	// lulu.print()

	//

	stevePointer := &steve
	stevePointer.updateName("Joey")
	steve.print()

	// Lulu is using a shortcut where go knows lulu is type of *Person / pointer to a Person type
	lulu.updateName("KogMaw")
	lulu.print()

}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

// Why are we using pointer here:
// Simple: the FirstName value is saved onto another address space than the Person struct
// p in updateName is a copy
// So when we update the FirstName's value we want to update the Person struct address space with a pointer

/////// This now becomes a type description - aka pointer to a person``
func (pointerToPerson *Person) updateName(newFirstName string) {
	// * operator here means we want to turn the pointerToPerson into a value of type Person
	(*pointerToPerson).FirstName = newFirstName
}

// Does not update FirstName!
// func (p Person) updateName(newFirstName string) {
// 	p.FirstName = newFirstName
// }

// Does update FirstName but I don't really know why!
// func (p *Person) updateName(newFirstName string) {
// 	p.FirstName = newFirstName
// }
