package main

import "fmt"

func main() {
	ids := []int{23, 12, 34, 37, 18}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	// Range with map
	data := map[string]string{"Jesus": "jesus@mail.com", "Alis": "alis@mail.com"}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range data {
		fmt.Println("Name: " + k)
	}
}
