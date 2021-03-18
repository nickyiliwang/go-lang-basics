package main

import (
	"log"
	"time"
)

// variable shadowing is when package level variable declaration and function scope var declation replacement because of scope

var s = "seven"

// createing an user verbose way!
// super tedious!
// var firstName string
// var lastName string
// var phoneNumber int
// var age int
// var birthDate time.Time

// cool way with struct
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	// GO LANG SPECIAL TYPE FOR DATES
	BirthDate time.Time
}

func main() {
	user := User{
		FirstName:   "Nick",
		LastName:    "Wang",
		PhoneNumber: "123-123-123-123-123",
	}

	log.Println(user.FirstName, user.LastName, user.Age)

	// var s2 = "six"
	// log.Println(s)
	// log.Println("s2 is:", s2)
	// gettingShadowed("xxx")
}

func gettingShadowed(s3 string) (string, string) {
	log.Println(s, "world")
	return s3, "world"

}
