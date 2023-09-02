package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// inc := 0

	routines := 50

	var wg sync.WaitGroup
	var inc int32

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			atomic.AddInt32(&inc, 1)
			fmt.Println(atomic.LoadInt32(&inc))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc)
}

/*
Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic
*/
