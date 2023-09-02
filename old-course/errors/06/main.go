package main

import (
	"fmt"

	Dog "github.com/sagivO2150/Hands-on-exercises/old-course/errors/06/dog"
)

// main defines the age of the dog in human years and sends the dog to the dog package which returns the age in dog years.
func main() {
	// BlAbLa := 7
	mydog := 7
	fmt.Println("my dogs age is- ", Dog.YEARS(mydog))

}

/*
Hands-on exercise #1
Create a dog package. The dog package should have an exported func “Years” which takes
human years and turns them into dog years (1 human year = 7 dog years). Document your code with comments. Use this code in func main.
● run your program and make sure it works
● run a local server with godoc and look at your documentation.
*/
