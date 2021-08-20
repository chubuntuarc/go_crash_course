package main

import "fmt"

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) Greeting() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spousLastname string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spousLastname
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Alis", lastName: "Test", city: "San Francisco", gender: "female", age: 30}
	// Alternative
	person2 := Person{"John", "Test", "San Francisco", "male", 30}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	fmt.Println(person1.Greeting())

	person1.hasBirthday()
	fmt.Println(person1.age)

	person1.getMarried("Arciniega")
	fmt.Println(person1.Greeting())

}
