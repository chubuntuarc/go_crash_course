package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assing kv
	emails["test"] = "test@mail.com"
	emails["jesus"] = "jesus@mail.com"
	emails["alis"] = "alis@mail.com"

	fmt.Println(emails)
	fmt.Println(emails["test"])

	// Delete from map
	delete(emails, "test")
	fmt.Println(emails)

	//Declare map and add kv
	data := map[string]string{"Jesus": "jesus@mail.com", "Alis": "alis@mail.com"}
	fmt.Println(data)
}
