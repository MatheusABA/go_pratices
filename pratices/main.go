package main

import (
	"fmt"
)

func main() {

	// Salute
	// name := "Matheus"
	// salute(name)

	// Doing integer division
	numerator := 10
	denominator := 0
	result, remainder, error := intDivision(numerator, denominator)
	// One way to check conditions is with if-else
	if error != nil {
		fmt.Println("Error:", error)
	} else if remainder == 0 {
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	} else {
		fmt.Printf("The result of %d divided by %d is %d and the remainder is %d\n", numerator, denominator, result, remainder)
	}
	// another way to do this is using switch
	switch {
	case error != nil:
		fmt.Println("Error:", error)
	case remainder == 0:
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	default:
		fmt.Printf("The result of %d divided by %d is %d and the remainder is %d\n", numerator, denominator, result, remainder)
	}
}

func salute(name string) {

	fmt.Printf("Hello %s, it was so a long time\n", name)

	welcomeAgain(name)

}

func welcomeAgain(name string) {

	fmt.Printf("Welcome again, %s!\n", name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var error error
	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator cannot be zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator

	// nil is used to indicate that there is no error during the execution of the program
	// if error return nil, it means that there is no error
	// return the result of the division, remainder an nil
	return result, remainder, error
}
