package main

import (
	"fmt"
	"strings"
)

func main() {

	String()

	concatenateString()
}

func String() {
	// Dealing with strings
	myString := "Matheus"
	index := myString[0]
	fmt.Printf("Hex: %x, Char: %c, Decimal: %d, Octal: %o, Type: %T\n", index, index, index, index, index) // Will return the value based on ASCII table	-> x = hex, c -> char, d -> decimal, %o -> octal
	for k, v := range myString {
		fmt.Println(k, v)
	}

	// Now with runes -> int32
	myRune := []rune("Matheus")
	indexRune := myRune[0]
	fmt.Printf("Hex: %x, Char: %c, Decimal: %d, Octal: %o, Type: %T\n", indexRune, indexRune, indexRune, indexRune, indexRune) // Will return the value based on Unicode table	-> x = hex, c -> char, d -> decimal, %o -> octal
	for k, v := range myRune {
		fmt.Println(k, v)
	}
}

func concatenateString() {
	// Concatenating strings
	newString := []string{"M", "a", "t", "h", "e", "u", "s"}
	concatString := ""
	for i := range newString {
		concatString += newString[i]
	}
	fmt.Println("Concatenated String:", concatString) // Worst case because it creates a new string every time you concatenate, so it is better to use strings.Builder or bytes.Buffer for large strings

	//  Using strings.Builder
	newStringBuilder := []string{"M", "a", "t", "h", "e", "u", "s"}
	stringBuilder := strings.Builder{}
	for i := range newStringBuilder {
		stringBuilder.WriteString(newStringBuilder[i])
	}
	stringBuilderString := stringBuilder.String()                         // Only now the string is created
	fmt.Println("Concatenated String with Builder:", stringBuilderString) // Better performance for large strings
}
