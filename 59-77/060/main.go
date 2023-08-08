package main

import "fmt"

func main() {
	fmt.Println("the odd funcs will run at the end")

	// defer fmt.Println("1")
	// fmt.Println("2")

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			defer fmt.Println(i)
		} else {
			fmt.Println(i)
		}

	}

}

/*
Hands-on exercise #60 - defer func
● “defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run
*/
