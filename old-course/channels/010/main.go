package main

import (
	"fmt"
	"sync"
)

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, val)
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int, len(data))
	var wg sync.WaitGroup

	// Start consumers
	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	// Send data to the channel
	for _, val := range data {
		ch <- val
	}
	close(ch)

	// Wait for consumers to finish
	wg.Wait()
}
