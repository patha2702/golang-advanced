package main

import (
	"fmt"
	"time"
)


// Anonymous functions:
// - These are functions that are declared without names.
// - No name, hence cannot be used later
// - These functions can also be called using go routines.


func main() {
	go func() {
		fmt.Println("Hello World")
	}()

	time.Sleep(1 * time.Second)
}

