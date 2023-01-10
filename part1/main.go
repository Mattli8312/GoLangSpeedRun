/*
* Go programs are organized in package: collection of source files grouped in a single directory
* in this case, main is the name of the packege stored in the GOLANGRUNTHROUGH directory
* Package name doesn't need to be the same name as the file name
 */
package main

/*
 * standard library: similar to stdio for c++
 */
import (
	"fmt"
	"time"
)

//Compiling: go build {file name}
//Compiling and running: go run {file name}

/*
* All files in the same directory belong to the same package:
* we created a file called displayTime in showtime.go
 */

/*
* Remember: When linking a github repo with your local system, make sure to create the main module first:
* Move to your main repo for your directory, i.e. GOLANGRUNTHROUGH and type go mod init {github url}
 */

// Symantic for multiple return values
func do_something() (int, bool) {
	return 6, false
}

func main() {

	fmt.Println("Hello World!")
	displayTime()

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
