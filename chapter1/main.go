/*
* Go programs are organized in package: collection of source files grouped in a single directory
* in this case, main is the name of the packege stored in the GOLANGRUNTHROUGH directory
* Package name doesn't need to be the same name as the file name
 */
package main

/*
 * standard library: similar to stdio for c++
 */
import "fmt"

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

func main() {
	fmt.Println("Hello World!")
	displayTime()
}
