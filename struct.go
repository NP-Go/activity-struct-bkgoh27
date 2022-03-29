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
		fName:            "Annakin2",
		lName:            "Skywalker2",
		age:              502,
		subscriber:       true,
		homeAddress:      "Death Star2",
		phone:            1234567,
		creditAvailable:  10000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	// customerbase := []customer{customer1, customer2}
	fmt.Println(customer1.fName, customer2.lName)
}
