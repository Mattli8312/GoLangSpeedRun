package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	FirstName string
	LastName  string
}

type Person struct {
	FirstName string
	LastName  string
	Details   struct {
		Height int
		Weight float32
	}
}

func chapter9() {
	// Encoding and Decoding Data Using JSON

	fmt.Println("Chapter 9")

	// Json supports the following data types
	/**
	Object
	String
	Boolean
	Number
	Array
	null
	*/

	// Be sure to capitalize the first character of each field in the people struct
	// if the field name starts with a lowercase letter, it won't be exported beyong the current package

	var person people
	jsonString := `{ "firstName" : "Matt" , "lastName" : "Lian"}`
	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.FirstName, person.LastName)
	} else {
		fmt.Println("Error: ", err)
	}

	// Decoding JSON to arrays
	var peoples []people
	jsonString = `[ {
		"firstName" : "John", "lastName" : "Snow"
	}, {
		"firstName" : "Nate", "lastName" : "Shoe"
	}]`
	err = json.Unmarshal([]byte(jsonString), &peoples)
	if err == nil {
		for _, v := range peoples {
			fmt.Println(v)
		}
	} else {
		fmt.Println("Error ", err)
	}

	// Decoding embedded objects
	var pops []Person
	jsonString = `[
		{
			"firstName" : "Donald",
			"lastName" : "Trump",
			"details" : {
				"height" : 175,
				"weight" : 70.0
			}
		},
		{
			"firstName" : "Rachel",
			"lastName" : "Adams",
			"details" : {
				"height" : 170,
				"weight" : 60.2
			}
		}
	]`
	err = json.Unmarshal([]byte(jsonString), &pops)
	if err == nil {
		for _, v := range pops {
			fmt.Println(v.FirstName, v.LastName, v.Details.Height, v.Details.Weight)
		}
	} else {
		fmt.Println(err)
	}
}
