package main

import "fmt"

func main() {
	//Using var
	var age int = 30
	var isCool = true
	//shorthand - only inside functions
	name, email := "Test", "test@mail.com"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", isCool)
}
