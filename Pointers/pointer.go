package main

// https://tour.golang.org/methods/4

import "fmt"

func main() {
	var coolCar string
	coolCar = "Yellow Buggy"

	fmt.Println("my coolCar is a", coolCar)
	upgradeMyCoolCarWithPointer(&coolCar)

	fmt.Println("now I have a cooler car", coolCar)

}

func upgradeMyCoolCarWithPointer(s *string) {
	fmt.Println("s is set to", s)
	newCar := "Silver Tesla Roadster"
	*s = newCar
}
