package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for z := 0; z < 10; z++ {
				c <- z
			}

		}()

	}

	for r := 0; r < 100; r++ {
		fmt.Println(r, <-c)
	}
}

/*
Hands-on exercise #7
● write a program that
○ launches 10 goroutines
■ each goroutine adds 10 numbers to a channel
○ pull the numbers off the channel and print them
*/
