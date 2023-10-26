package main

import (
	"fmt"
	"time"
)


// Select statement:
// - The select statement in Go looks like a switch statement
//   but for channels.
// - The select statement lets a go routine wait on multiple communication
//   operations.
// - In select, each of the case statement waits fo a send or receive operation
//   operation from channel
// - select blocks, until any of the case statements are ready
// - If multiple case statements are ready, then it selects one at random and proceeds.
// - select statement lets a go routine wait on multiple communication operations.
// - Default case will be executed if no send or receive operation is ready on any of the case statements.
// - Defualt case makes the select statement non blocking, i.e. if block operation is there, and no operation
//   is ready, then default case will be executed.
// - Powerful tool for synchronization and concurrency.


// Switch vs Select:
// - Non blocking    |     Blocking (can be non blocking with default case)
// - Deterministic   |     Non Deterministic


func chOne(ch chan int){
	ch <- 1
}

func chTwo(ch chan int) {
	ch <- 2
}


func main() {
	ch := make(chan int, 5)
	ch2 := make(chan int, 5)

	go chOne(ch)
	go chTwo(ch2)

	select {
		case val1 := <- ch:
			fmt.Println(val1)
		case val2 := <- ch2:
			fmt.Println(val2)
		default:
			fmt.Println("No operation is ready.")
	}
	time.Sleep(2 * time.Second)
}
