package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

/*
Hands-on exercise #5
We are going to learn about testing next. For this exercise, take a moment and see how much
you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s article on testing. See if you can get a basic example of testing working.
*/
