package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)
	fmt.Printf("the value of x is - %v \nthe value of y is - %v \n", x, y)

	/*
		if x <= 4 && y <= 4 {
			fmt.Printf("so they are both smaller or equal to 4\n")
		} else if x >= 6 && y >= 6 {
			fmt.Printf("so they are both grater then 6\n")
		} else if x >= 4 && x <= 6 {
			fmt.Printf("so x is between 4 to 6\n")
		} else if y != 5 {
			fmt.Printf("and so y is not equal to 5\n")
		} else {
			fmt.Printf("non of those conditsion were met\n")
		}
	*/
	switch {
	case x <= 4 && y <= 4:
		fmt.Printf("so they are both smaller or equal to 4\n")
	case x >= 6 && y >= 6:
		fmt.Printf("so they are both grater then 6\n")
	case x >= 4 && x <= 6:
		fmt.Printf("so x is between 4 to 6\n")
	case y != 5:
		fmt.Printf("and so y is not equal to 5\n")
	default:
		fmt.Printf("non of those conditsion were met\n")

	}

}
