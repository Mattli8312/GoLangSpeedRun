package main

import (
	"fmt"
	"time"
)

func displayTime() {
	fmt.Println("Time: ", time.Now())
}

// Defining functions with multiple parameters
func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

// passing arguments by value
func swap(a, b int) {
	a, b = b, a
	// just like python!
	// However, this won't swap the values in main
}

// Passing arguments by pointer
func swap_by_pointer(a, b *int) {
	*a, *b = *b, *a
}

// returning value
func addNum(num1, num2 int) int { // return int
	return num1 + num2
}

// returning multiple values
func countOddEven(s string) (int, int) {
	odd, even := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return odd, even
}

// naming return values
// instead of return sum, you can just specify the return values and type return
func addNum2(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

// variadic functions:
// functions that take variable number of arguments. most common is fmt.Println()
func addNumsVar(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// implementing closure using anonymous functions

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}
