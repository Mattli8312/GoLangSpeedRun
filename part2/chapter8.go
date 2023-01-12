package main

import (
	"fmt"
	"sort"
)

// Creating a map of structs
type dob struct {
	day   int
	month int
	year  int
}

type person struct {
	name  string
	email string
	dob   dob
}

func chapter8() {
	fmt.Println("Chapter 8")
	// Establishing relationships using maps

	// A go map has the following syntax:
	// map[keyType] valueType
	var heights map[string]int
	heights = make(map[string]int)
	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jan"] = 175

	fmt.Println(heights["Peter"])
	fmt.Println(heights["Joan"])
	fmt.Println(heights["Jan"])

	// Intializing a map with a map literal
	widths := map[string]int{
		"Peter": 32,
		"Jan":   30,
		"Joan":  28, // <-- Note the comma here
	}

	// Checking the existenct of a key
	fmt.Println(widths["Jim"])

	// However, this doesn't tell you if the value of the key is zero or if the key
	// Simply doesn't exists, hence it's better to do this

	// Notice below, you can chain expressions in your if statement
	if v, ok := heights["Jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}
	fruits := map[string]int{
		"orange":     20,
		"apple":      24,
		"guava":      43,
		"watermelon": 10,
	}

	//delete an item
	delete(fruits, "guava")

	//display the map
	fmt.Println(fruits)

	// Getting the number of items in a map
	fmt.Println(len(heights))

	// Iterating over a mpa
	for k, v := range heights {
		fmt.Println(k, v)
	}

	keys := []string{}
	// Getting all the keys in a map
	for k := range heights {
		keys = append(keys, k)
	}
	//Note: the order of the keys retrieved is not guaranteed

	//Setting iteration order in a map
	sort.Strings(keys)
	fmt.Println(keys)

	// Iterate through map
	for _, k := range keys {
		fmt.Println(heights[k])
	}

	// to store a collection of person struct, you can declare a map of int key and value of type person
	members := make(map[int]person)

	members[1] = person{
		name:  "Mary SMith",
		email: "marysmith@example.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}

	members[4] = person{
		name:  "John Stu",
		email: "johnstu@example.com",
		dob: dob{
			day:   12,
			month: 4,
			year:  1994,
		},
	}

	for k, v := range members {
		fmt.Println(k, v)
	}
}
