package main

import (
	"fmt"
	"time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) { //creating a function that recives a function to test (can be any func)
	start := time.Now()
	fn() //f represemts the function that we will pass in to test how long it takes to run
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed) // prints the amount of time passed since the func started
}

// Function to be wrapped
func MyFunction() { //this will be the function that we will pass to the TimedFunction
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Println("MyFunction completed")
}

func main() {
	// Call the wrapped function
	TimedFunction(MyFunction) //here we are "summening" the time testing function and passing in to it the function that we want to test how much time it took it to run
}
