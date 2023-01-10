package main

import (
	"fmt"
)

// Symantic for multiple return values
func do_something() (int, bool) {
	return 6, false
}

func main() {
	fmt.Println("Making Decision")

	num := 6
	//Expressions
	fmt.Println(num == 0)

	fmt.Println(num != 0)

	//Control statements
	if num != 0 {
		fmt.Println("num != 0")
	} else {
		fmt.Println("num == 0")
	}

	if v, err := do_something(); !err {
		fmt.Println(v)
	} else {
		fmt.Println("Error")
	}
	dayofweek := "Monday"
	switch num {
	case 1:
		dayofweek = "Tuesday"
	default:
		dayofweek = "Friday"
	}
	println("Day of the week, ", dayofweek)

	//Switch with control statements
	score := 79
	grade := ""
	switch {
	case score < 50:
		grade = "F"
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	default:
		grade = "A"
	}
	fmt.Println("Your grade: ", grade)
}
