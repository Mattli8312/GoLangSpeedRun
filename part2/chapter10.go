package main

import (
	"fmt"
	"math"
)

// Defining an interface
type DigitsCounter interface {

	// Only has 1 method at this point: returns two results int and int
	// Any type that implements the countoddeven method will also implement the digitscounter interface

	CountOddEven() (int, int)
}

// Implementing an interface
type DigitString string

// DigitString implements DigitsCounter
// Notice we pass in the receiver to the beginning of the function

func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

// Looking at how you may use interfaces

type circle struct {
	radius float64
	name   string
}

type square struct {
	length float64
	name   string
}

// You also need the area of these shapes: these are like virtual methods in c++
type shape interface {
	Area() float64 // both the square and the circle will implement the shape interface
}

// circle implements shape
func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// square implements shape
func (s square) Area() float64 {
	return s.length * s.length
}

// function for printing area of all shapes
func calculateArea(listOfShapes []shape) {
	for _, shape := range listOfShapes {
		fmt.Println("area of shape: ", shape.Area())
	}
}

// Although the shape interface has a single method signature, Area()
// Types that implement this interface are not restricted to only implementing the Area()

func (c circle) Circumference() float64 {
	// Additional method for circle
	return 2 * math.Pi * c.radius
}

// Using the stringer interface
// Interfaces may still seem abstract to you at this point,

// person struct:
type Person_ struct {
	Firstname string
	Lastname  string
	Age       int
}

// The sringer interface is a type that describes itself as a string

type Stringer interface {
	String() string
}

func (p Person_) String() string {
	// Here, we are overriding the default behavior for the struct's String() function.
	// Now when you print out the value of me,
	return fmt.Sprintf("%v %v (%d years old)", p.Firstname, p.Lastname, p.Age)
}

// using an empty interface
// Useful when you want to handle data of unknown types
func doSomething(i interface{}) {
	fmt.Println(i)
}

func chapter10() {
	fmt.Println("Chapter 10")

	// Interface: Defines the behavior of an object, specifying methods that it needs to implement
	s := DigitString("1234567890")
	fmt.Println(s.CountOddEven()) // 5, 5

	// Call instances of these objects
	c := circle{radius: 4.4, name: "circle"}
	fmt.Println(c.Area())
	sq := square{length: 3.5, name: "square"}
	fmt.Println(sq.Area())

	// But this is too tediuos. Because each shape implements the Shape interface, they all have the Area method
	// Instead, we should iterate through all of the shapes and call the Area function on these shapes
	shapes_list := []shape{c, sq}
	calculateArea(shapes_list)

	// Adding methods to a type that doesn't satisfy an interface
	fmt.Println(c.Circumference())

	var str_ [2]string
	str_[0] = "Hello"
	str_[1] = "World"
	fmt.Println(str_)

	me := Person_{"Matthew", "Lian", 21}
	fmt.Println(me)

	// empty interfaces
	doSomething("Hi!")          // string
	doSomething(3.15)           // Float
	doSomething([]int{1, 2, 3}) // Array
}
