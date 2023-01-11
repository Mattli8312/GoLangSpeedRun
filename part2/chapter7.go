package main

import (
	"fmt"
	"math"
)

// The following is a struct called point
type point struct {
	x float32
	y float32
	z float32
}

// Defining methods in Structs
// Notice how it has the p point receiver: the parameter before the function name is the
func (p point) length() float64 {
	return math.Sqrt((math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2) + math.Pow(float64(p.z), 2)))
}

func createNewStruct() *point {
	p := point{x: 1.2, y: 2.4, z: 4.8}
	return &p
}

// Comparing Structs
func comparing_structs() {
	var pt1 = point{x: 1.1, y: 2.2, z: 3.3}
	var pt2 = point{x: 1.1, y: 2.2, z: 3.3}
	fmt.Println(pt1 == pt2)
}

// Copying Strucs
func copying_structs() {
	pt4 := createNewStruct()
	pt5 := pt4

	fmt.Println(pt4, pt5)
	pt5.x = 5.5
	fmt.Println(pt4, pt5)

	// These two pointers point to the same struct, but what if we want a deep copy?
	// If we want an independent copy, we have to use the * character like this
	pt6 := *pt4
	fmt.Println(pt4, pt6)
	pt6.x = 12.3
	fmt.Println(pt4, pt6)
}

func chapter7() {
	/*Defining the Blueprints of Your Data Using structs*/

	fmt.Println("Chapter 7")

	// Creating the object
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 6.2

	// struct literal
	pt2 := point{x: 5.6, y: 1.3, z: 2.2} // Like python
	fmt.Println(pt1, pt2)

	pt3 := createNewStruct()
	fmt.Println(pt3)

	// Notice, it will print out &{1.2, 2.4, 4.8}
	fmt.Println(pt3.x)
	// This will automatically dereference pt3 and print out the member

	// Calling member
	fmt.Println(pt3.length())

	// Comparing Structs
	comparing_structs()

	// Copying Structs
	copying_structs()
}
