package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)

	}()

	for i := range c {
		fmt.Println(i)
	}

}

/*
Hands-on exercise #6 ● write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
*/
