package main 

import (
	"fmt"
	"sync"
)


// Wait groups:
// - We know that the main go routine exits, even when other go routines are executing their task.
// - To solve that we used sleep method, but it was not an elegant solution.
// - Hence, we have wait groups.
// - To wait for multiple go routines to finish, we can use wait groups.
// - It is a synchronization primitive used to coordinate the execution of multiple goroutines and wait for them
//   to finish their tasks before proceeding with further execution.
// - It implements an internal counter to keep track of multiple go routines.


func main() {
	// Defining the wait group
	var wg sync.WaitGroup

	for i:=1; i<5; i++ {
		wg.Add(1) // incrementing the internal counter by number of go routines we are adding
		go func(i int) {
			fmt.Printf("Go routine %d", i)
			defer wg.Done() // This method will run when the function it is defined in, is executed.
			// It will decrement the internal counter
		}(i)
	}
	wg.Wait()
	// This will block the main go routine until counter becomes 0
	// Wait for all go routines to finish.
	fmt.Println("All go routines executed")
}
