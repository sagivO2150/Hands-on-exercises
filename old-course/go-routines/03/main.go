package main

import (
	"fmt"
	"sync"
)

func main() {
	inc := 0

	routines := 50

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			mu.Lock()
			v := inc
			v++
			inc = v
			fmt.Println(inc)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc)
}

/*
Hands-on exercise #3
● Using goroutines, create an incrementer program
○ have a variable to hold the incrementer value
○ launch a bunch of goroutines
■ each goroutine should
● read the incrementer value
○ store it in a new variable
● yield the processor with runtime.Gosched()
● increment the new variable
● write the value in the new variable back to the incrementer
variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
*/
