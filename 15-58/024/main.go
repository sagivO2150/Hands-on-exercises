package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(350)
	fmt.Printf("the value of x is - %v and it is ", x)

	/*
		if x <= 100 {
			fmt.Printf("between 0 and 100\n")
		} else if x > 100 && x <= 200 {
			fmt.Printf("between 101 and 200\n")
		} else if x > 200 && x <= 250 {
			fmt.Printf("between 201 and 250\n")
		}

	*/
	switch {
	case x <= 100:
		fmt.Printf("between 0 and 100\n")
	case x > 100 && x <= 200:
		fmt.Printf("between 101 and 200\n")
	case x > 200 && x <= 250:
		fmt.Printf("between 201 and 250\n")
	default:
		fmt.Printf("the value of x is grater then 250\n")
	}

}
