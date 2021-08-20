package main

import "fmt"

func main() {
	// Arrays
	var fruitsArray [2]string

	// Assign values
	fruitsArray[0] = "Apple"
	fruitsArray[1] = "Orange"

	// Declare and assing
	fruitData := [2]string{"Apple", "Orange"}

	fmt.Println(fruitsArray)
	fmt.Println(fruitData)

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

}
