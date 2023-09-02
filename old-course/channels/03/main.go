package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println("this is a recived number from channel c-", i)
	}

}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}

	}()
	return c
}

/*
Hands-on exercise #3
● Starting with this code, pull the values off the channel using a for range loop
*/
