// Go Runtime Scheduler:

// What happens when you start a go program?
// - Go program will launch operating system threads equal to available number of 
//   logical CPU's.
// - These threads are OS threads, completely managed by OS/ kernel.
//   i.e. creating, blocking, and scheduling them on cpu cores.
// - Find number of logical processors => runtime.NumCPU() method from runtime package
// - Logical cores = no of physical cores * no of threads that can run on each core(hardware threads)
 
package main

import "runtime"
import "fmt"

func main() {
	fmt.Println(runtime.NumCPU())
}
