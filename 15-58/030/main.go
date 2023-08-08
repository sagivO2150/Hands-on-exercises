package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 1; i < 43; i++ {

		// fmt.Printf("this program ran %v times\n", i)

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("this program ran %v times\t x equals to - %v\n", i, x)
		case 1:
			fmt.Printf("this program ran %v times\t x equals to - %v\n", i, x)
		case 2:
			fmt.Printf("this program ran %v times\t x equals to - %v\n", i, x)
		case 3:
			fmt.Printf("this program ran %v times\t x equals to - %v\n", i, x)
		case 4:
			fmt.Printf("this program ran %v times\t x equals to - %v\n", i, x)
		default:
			fmt.Printf("this program ran %v times\t x is above 4", i)
		}
	}
}
