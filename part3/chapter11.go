package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// 100 ms delay
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

var balance int64

// mutexes
var mutex = &sync.Mutex{}

func credit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance += 100
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After crediting, balance is ", balance)
		mutex.Unlock() // Need this to prevent from deadlocking
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance -= 100
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After debiting, balance is", balance)
		mutex.Unlock()
	}
}

func atomic_credit() {
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&balance, 30)
		fmt.Println("After atomic credit, balance is ", balance)
	}
}

func chapter11() {
	fmt.Println("Chapter 11: Threading Using Goroutines")
	// In go, goroutines are functions that run concurrently with other functions
	// When you use go-routines, your program will be more responsive

	// You might use go-routines when performing tasks that deal with different input sources.
	// i.e. networks, backend services, etc.

	say("Hello", 3)
	say("World", 2)

	// to run the function as a go routine, use the go keyword
	// go say("Hello", 3)
	// go say("World", 2)
	// fmt.Scanln()

	// The reason we need the fmt.Scanln() is because each time a goroutine is called, the control is immediately
	// returned to the claling statement
	// Without scanln() function to wait for user input,
	// the program automatically terminates after the second goroutine is called.

	/*Using Goroutines with Shared Resources*/
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	go atomic_credit()
	fmt.Scanln()

	// Using atomic counters for modifying shared resources

}
