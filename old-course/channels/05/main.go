package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v = <-c
	fmt.Println(v)

}

/*
Hands-on exercise #5
● Show the comma ok idiom starting with this code.
*/
