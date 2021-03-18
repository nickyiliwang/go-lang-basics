package main

// go tool cover -help
// go test -coverprofile=c.out
// go tool cover -html=c.out

import (
	"math"
	"testing"
)

// Doing a function for each test seems inefficient
// So we are trying this

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, float32(math.Inf(0)), true},
	{"expect-5", 25.0, 5.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		// tt.isErr is true, we should expect an error
		if tt.isErr {
			// if no err present
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// Doing a function for each test seems inefficient

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(0, 1.0)
// 	if err == nil {
// 		t.Error("x is 0 and we should not divide by 0")
// 	}
// }
