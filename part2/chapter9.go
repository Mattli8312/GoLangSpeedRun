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

type Rates struct {
	// struct field tags (string literal placed after each field in a struct)
	// Sometimes keys in JSON string can't be mapped directly to members of your struct in GO
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

type FormalName struct {
	Firstname string
	Lastname  string
}

type Address struct {
	Line1 string
	Line2 string
	Line3 string
}

type Customer struct {
	Name    FormalName
	Email   string
	Address Address
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

	jsonString = `{ "base currency" : "EUR", "destination currency" : "USD" } `
	var rates Rates
	err = json.Unmarshal([]byte(jsonString), &rates)
	if err == nil {
		fmt.Println(rates.Base)
		fmt.Println(rates.Symbol)
	} else {
		fmt.Println(err)
	}

	daniel := Customer{
		Name:  FormalName{Firstname: "Daniel", Lastname: "Sald"},
		Email: "123@gmail.com",
		Address: Address{
			Line1: "123",
			Line2: "456",
			Line3: "789",
		},
	}

	res, errRes := json.Marshal(daniel)
	if errRes == nil {
		fmt.Println(string(res))
	}
}
