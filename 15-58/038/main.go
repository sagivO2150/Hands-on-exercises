package main

import (
	"fmt"
	"math/rand"
)

func main() {

	c := 1

	for i := 0; i < 100; i++ {

		if x := rand.Intn(5); x == 3 {
			fmt.Printf("x IS 3 \t it appered in iteration number - %v \t and it appered %v times in total\n", i, c)
			c++
		}

		// } else {
		// 	fmt.Println("x IS NOT 3")
		// }
		// fmt.Printf("this loop ran %v times\n", i)

	}

}
