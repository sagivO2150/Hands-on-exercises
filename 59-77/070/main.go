package main

import "fmt"

func first() int {
	num := second()
	return num
}

func second() int {
	return 42
}

func outer() func() int {
	return func() int {
		return 43
	}
}

func main() {
	x := first()
	fmt.Println(x)

	fmt.Println(outer()())
}

/*
Hands-on exercise #70 - func return
● Create a func
○ which returns a func
■ which returns 42
● assign the returned func to a variable
● call the returned func
● print
*/
