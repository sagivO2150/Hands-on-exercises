package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin routines", runtime.NumGoroutine())
	fmt.Println("")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("first go routine")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		fmt.Println("second go routine")
		wg.Done()
	}()

	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid routines", runtime.NumGoroutine())
	fmt.Println("")

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end routines", runtime.NumGoroutine())
	fmt.Println("")

}

/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/
