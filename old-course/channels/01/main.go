package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	// go func() {
	// 	c <- 42
	// }()

	c <- 42

	fmt.Println(<-c)

}

/*
Hands-on exercise #1
● get this code working:
○ using func literal, aka, anonymous self-executing func ■ solution: https://play.golang.org/p/SHr3lpX4so
○ a buffered channel
*/
