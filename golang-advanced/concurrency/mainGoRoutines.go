package main

import (
	"fmt"
	"time"
)


// Main Go Routine:
// - The main function from the main package starts the main goroutine.
// - All goroutines, including the main one, are started from the main goroutine,
//   and they can create other goroutines.
// - The main goroutine runs the main program.
// - Goroutines exit when their associated function returns.
// - There is no parent-child relationship between goroutines; they are independent.
// - When you start a go routine, it executes alongside all other go routines.
// - When the main goroutine exits (e.g., when the 'main()' function returns),
//   it indicates that the program has finished, and all other goroutines are
//   abruptly terminated, even if their functions have not returned.


// Goroutine Relationships in Go:
// - In Go, we can create goroutines to run tasks concurrently.
// - A function that starts a goroutine is like a "task creator."
// - The creator function uses 'go' to start a new goroutine.
// - The created goroutine is independent and runs its task concurrently.
// - There's no parent-child hierarchy; both are peers.
// - They can communicate but don't control each other's execution.


func start() {
	go process()
	fmt.Println("In start")
}


func process() {
	fmt.Println("In process")
}

func main() {
	go start()
	time.Sleep(3 * time.Second)
}
