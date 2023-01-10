package main

import (
	"fmt"
	"time"
)

/**
Four basic data-types
Basic: strings, numbers, booleans
Aggregate: arrays and structs
Reference: pointers, slices, functions, and channels
Interface: An interface is a collection of method signatures
*/

// Declaring Never-Changing Constants
// If you set a variable outside a function, the compiler won't complain
// However if you place it inside a function, it will
const publisher = "Wiley"

func main() {
	fmt.Println("Data Types")

	var numx = 5
	_ = numx //Compiler is now happy!

	//Using the var keyword: (like javascript)
	var num1 = 5     //type inferred
	var num2 int     //type explicit: when printed out, it will display 0
	var num3 float32 // prints 0
	var num4 string  // prints an empty string
	var raining bool // prints boolean variable
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3, num4, raining)

	//You can also explicitly declare and initialize a value at the same time
	var rates float32 = 4.5
	fmt.Println(rates)

	//Using the short variable declaration operator
	//Also can streamline variable declarations like python
	firstName, lastName := "Wei-Meng", "Lee"
	// inferred and declared
	fmt.Println("firstName: ", firstName, " lastName: ", lastName)

	//Another way to do the above is
	var (
		first string = "Matthew"
		last  string = "Lian"
		age   int    = 21
	)
	fmt.Println("My name is ", first, last, " and I am ", age, " years old")

	//The := operator can only be used for declaring and initializing variables inside functions

	//Determining the data type:
	start := time.Now()
	fmt.Printf("Data type: %T\n", start) //time.Time

}
