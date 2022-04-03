package main

import "fmt"

// declaration of struct
type customer struct {
	fName            string
	lName            string
	age              int
	subscriber       bool
	homeAddress      string
	phone            int
	creditAvailable  float32
	currentCartCost  float32
	currentOrderCost float32
}

func main() {
	// create the object
	customer1 := customer{
		fName:            "Annakin",
		lName:            "Skywalker",
		age:              50,
		subscriber:       true,
		homeAddress:      "Death Star",
		phone:            1234567,
		creditAvailable:  10000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	customer2 := customer{
		fName:            "Han",
		lName:            "Solo",
		age:              50,
		subscriber:       false,
		homeAddress:      "Tatooine",
		phone:            4321765,
		creditAvailable:  20000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	// customerbase := []customer{customer1, customer2}

	// access the info
	fmt.Println(customer1.fName, customer2.lName)

	fmt.Printf("%+v\n", customer1) // Print with Variable Name
	fmt.Printf("%v\n", customer2)  // Without Variable Name
}

// note: this is not databse nor sql
