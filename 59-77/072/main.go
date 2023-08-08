package main

import (
	"fmt"
	"math"
)

func makeThePower(powerOf float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		results := math.Pow(powerOf, c)
		// results := 1.0
		// for i := 0; i < powerOf; i++ {
		// 	// fmt.Printf("the loop ran with the base equals - %v to the power of -%v \n", results, base)
		// 	fmt.Printf("stage %v: \tis %v to the power of %v\n", i, results, base)
		// 	results = results * base
		// }
		return results
	}
}

func main() {
	calc := makeThePower(5)
	// fmt.Printf("and the finale results are - %v\n", calc(2))
	// fmt.Println(calc(2))
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
	fmt.Println(calc())
}

/*
Hands-on exercise #72 - closure
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/
