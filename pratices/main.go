package main

import "fmt"

type Person struct {
	name    string
	age     int
	gender  string
	address string
	Car     // This is an embedded struct
}

type Car struct {
	model string
	year  int
	owner string
	gasEngine
}

type gasEngine struct {
	kmpl   uint16
	liters uint16
}

type engine interface {
	kmsLeft() uint16
}

func (e gasEngine) kmsLeft() uint16 {
	return e.kmpl * e.liters
}

func canReachDestination(e engine, kms uint16) {
	if kms <= e.kmsLeft() {
		fmt.Println("You can reach your destination")
	} else {
		fmt.Println("You cannot reach your destination")
	}

}

func main() {

	maleUser := Person{
		name:    "John Doe",
		age:     30,
		gender:  "Male",
		address: "123 Main St",
		Car: Car{ // Initializing the embedded struct
			model: "Toyota Camry",
			year:  2020,
			owner: "John Doe",
		},
	}

	myEngine := gasEngine{
		15,
		20,
	}

	maleUser.Car.gasEngine = myEngine // Assigning the gasEngine to the embedded struct

	fmt.Println("User Information")
	fmt.Printf("Name: %s\n", maleUser.name)
	fmt.Printf("Age: %d\n", maleUser.age)
	fmt.Printf("Gender: %s\n", maleUser.gender)
	fmt.Printf("Address: %s\n", maleUser.address)
	// We can access the fields of the embedded struct directly without needing to specify the struct name, in case that Person struct has a field carInfo Car, we would need to use maleUser.Car.[field] instead
	fmt.Printf("Car Model: %s\n", maleUser.model)
	fmt.Printf("Car Year: %d\n", maleUser.year)
	fmt.Printf("Car Owner: %s\n", maleUser.owner)
	fmt.Printf("Car KMPL: %d\n", maleUser.kmpl)
	fmt.Printf("Car Liters: %d\n", maleUser.liters)

	fmt.Print("--------------------------------")
	fmt.Printf("\nKms left in the car: %d\n", maleUser.kmsLeft()) // We can access the method of the embedded struct directly
	fmt.Print("--------------------------------\n")
	canReachDestination(myEngine, 200) // We can pass the embedded struct to a function

}
