package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case <-q:
			return
		case i := <-c:
			fmt.Println(i)
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
	}()

	return c
}

/*
Hands-on exercise #4
● Starting with this code, pull the values off the channel using a select statement
*/
