package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	//y := 5
	/*
			if z := 2 * rand.Intn(x); z >= x {
				fmt.Printf("z was randomized with the value - %v which it is grater then x \n", z)
			} else {
				fmt.Printf("z was randomized with the value - %v which it is smaller then x\n", z)
			}

		z := 2 * rand.Intn(x)

		switch {
		case z < 42:
			fmt.Printf("z is - %v which is LESS then 42\n", z)
		case z >= 42:
			fmt.Printf("z is - %v which is EQUAL or GRATER then 42\n", z)
			fallthrough
		case z > 42:
			fmt.Printf("z is - %v which is ONLY GRATER then 42\n", z)
		default:
			fmt.Printf("this will never run\n")
		}
	*/
	z := 2 * rand.Intn(x)
	//i := 0
	for i := 0; i < z; i++ {
		if i%2 != 1 {
			continue
		}
		fmt.Printf("the value of i is - %v\n", i)
	}

}
