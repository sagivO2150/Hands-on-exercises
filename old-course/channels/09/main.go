package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, ch chan<- int) {
	for i := 0; i < 1; i++ {
		ch <- id*10 + i
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	close(ch)
}

func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	merge := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			out <- val
		}
		fmt.Println("merge done")
	}

	wg.Add(len(inputs))
	for _, input := range inputs {
		go merge(input)

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(1, ch1)
	go producer(2, ch2)

	merged := fanIn(ch1, ch2)

	// fmt.Println("before merge 2")
	for val := range merged {
		fmt.Println(val)
	}
}
